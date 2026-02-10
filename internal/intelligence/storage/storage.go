package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

var mu sync.Mutex

func ensureDir() (string, error) {
	dir := ".wells"
	err := os.MkdirAll(dir, 0755)
	return dir, err
}

func ReadJSON(file string, out interface{}) error {
	mu.Lock()
	defer mu.Unlock()

	dir, err := ensureDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, file)
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(data, out)
}

func WriteJSON(file string, in interface{}) error {
	mu.Lock()
	defer mu.Unlock()

	dir, err := ensureDir()
	if err != nil {
		return err
	}

	path := filepath.Join(dir, file)
	data, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
