package core

import (
	"bytes"
	"crypto/sha256"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"github.com/iancoleman/orderedmap"
	"github.com/pt-main/pack/lib/core"
)

// SaveState walks through the directory `where` recursively,
// computes SHA256 hash for each file, and returns an ordered map
// where key = absolute file path, value = hex-encoded hash.
func SaveState(where string) (*orderedmap.OrderedMap, error) {
	om := orderedmap.New()

	abs, err := filepath.Abs(where)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(abs)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, os.ErrInvalid
	}

	err = filepath.WalkDir(abs, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		hash := sha256.Sum256(data)
		om.Set(path, hash[:])
		return nil
	})

	if err != nil {
		return nil, err
	}

	return om, nil
}

func Changes(was *orderedmap.OrderedMap, where string) ([]string, error) {
	res := []string{}
	now, err := SaveState(where)
	if err != nil {
		return res, err
	}
	wasKeys := was.Keys()
	for _, key := range now.Keys() {
		if slices.Contains(wasKeys, key) {
			w, _ := was.Get(key)
			c, _ := now.Get(key)
			if !bytes.Equal(w.([]byte), c.([]byte)) {
				res = append(res, key)
			}
		} else {
			res = append(res, key)
		}
	}
	return res, nil
}

func StateAsPackCore(data *orderedmap.OrderedMap) ([]byte, error) {
	c := core.NewCore(data)
	return c.CreateFile()
}

func PackCoreAsState(data []byte) (*orderedmap.OrderedMap, error) {
	c := core.NewCore(nil)
	err := c.ReadFile(data)
	return c.Containers, err
}
