package core

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/iancoleman/orderedmap"
	"github.com/pt-main/pack/lib/core"
	"github.com/pt-main/run/tal/lang"
	"github.com/pt-main/run/tal/shared"
)

// FileHash computes SHA256 hash of a file in a streaming fashion.
func FileHash(path string) ([]byte, error) {
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

// SaveState hashes every regular file under the directory `where` and returns
// an ordered map of absolute path to hash, hashing files in parallel.
func SaveState(where string) (*orderedmap.OrderedMap, error) {
	const maxWorkers = 32

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
			h, err := FileHash(path)
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

	// sort for a deterministic order in the returned map
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].path < entries[j].path
	})

	om := orderedmap.New()
	for _, e := range entries {
		om.Set(e.path, e.hash)
	}
	return om, nil
}

// Changes returns the files under `where` that are new or whose hash differs
// from the previously saved state `was`.
func Changes(was *orderedmap.OrderedMap, where string) ([]string, error) {
	now, err := SaveState(where)
	if err != nil {
		return nil, err
	}

	oldHashes := make(map[string][]byte, len(was.Keys()))
	for _, k := range was.Keys() {
		v, _ := was.Get(k)
		oldHashes[k] = v.([]byte)
	}

	var res []string
	for _, k := range now.Keys() {
		oldHash, exists := oldHashes[k]
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

func Update() error {
	st, err := SaveState(".")
	if err != nil {
		return err
	}
	file, err := StateAsPackCore(st)
	if err != nil {
		return err
	}
	return Write(shared.TalFile, file)
}

func OpenF(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("Open: %v", err)
	}
	return string(data), nil
}

func Open(file string) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Open: %v", err)
	}
	return data, nil
}

func Write(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Write: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("Write: %v", err)
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("Write: %v", err)
	}
	return nil
}
