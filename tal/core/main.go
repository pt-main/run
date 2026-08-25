package core

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/iancoleman/orderedmap"
	"github.com/pt-main/pack/lib/core"
	"github.com/pt-main/run/tal/lang"
)

// fileHash computes SHA256 hash of a file in a streaming fashion.
func fileHash(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// SaveState walks through the directory `where` recursively,
// computes SHA256 hash for each file, and returns an ordered map
// where key = absolute file path, value = hex-encoded hash.
//
// Uses parallel workers and streaming reads for better performance.
func SaveState(where string) (*orderedmap.OrderedMap, error) {
	const maxWorkers = 32 // can be adjusted or set to runtime.NumCPU()

	abs, err := filepath.Abs(where)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(abs)
	if err != nil || !info.IsDir() {
		return nil, os.ErrInvalid
	}

	var files []string
	err = filepath.WalkDir(abs, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		info, err := os.Stat(path)
		if err != nil {
			return nil
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	type result struct {
		path string
		hash []byte
		err  error
	}
	results := make(chan result, len(files))
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxWorkers)

	for _, p := range files {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			h, err := fileHash(path)
			results <- result{path, h, err}
		}(p)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var entries []result
	for r := range results {
		if r.err != nil {
			return nil, r.err
		}
		entries = append(entries, r)
	}

	// Sort for deterministic order in orderedmap
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].path < entries[j].path
	})

	om := orderedmap.New()
	for _, e := range entries {
		om.Set(e.path, e.hash)
	}
	return om, nil
}

// Changes compares a previous state (was) with the current state of the directory `where`.
// It returns a list of absolute file paths that are either new or modified.
// The comparison is based on SHA256 hashes.
func Changes(was *orderedmap.OrderedMap, where string) ([]string, error) {
	now, err := SaveState(where)
	if err != nil {
		return nil, err
	}

	// Build a map for O(1) lookup of old hashes
	wasMap := make(map[string][]byte, len(was.Keys()))
	for _, k := range was.Keys() {
		v, _ := was.Get(k)
		wasMap[k] = v.([]byte)
	}

	var res []string
	for _, k := range now.Keys() {
		oldHash, exists := wasMap[k]
		if !exists {
			res = append(res, k)
			continue
		}
		newHash, _ := now.Get(k)
		if !bytes.Equal(oldHash, newHash.([]byte)) {
			res = append(res, k)
		}
	}
	return res, nil
}

func StateAsPackCore(data *orderedmap.OrderedMap) ([]byte, error) {
	c := core.NewCore(data)
	res, err := c.CreateFile()
	if err != nil {
		err = errors.New(lang.GetRealErrorReverse(err))
	}
	return res, err
}

func PackCoreAsState(data []byte) (*orderedmap.OrderedMap, error) {
	c := core.NewCore(nil)
	err := c.ReadFile(data)
	if err != nil {
		err = errors.New(lang.GetRealErrorReverse(err))
	}
	return c.Containers, err
}
