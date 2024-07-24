package storage

import (
	"os"
	"path/filepath"
)

const storageDir = "uploads"

func SaveFile(filename string, data []byte) error {
	if _, err := os.Stat(storageDir); os.IsNotExist(err) {
		os.Mkdir(storageDir, os.ModePerm)
	}

	path := filepath.Join(storageDir, filename)
	return os.WriteFile(path, data, 0644)
}

func LoadFile(filename string) ([]byte, error) {
	path := filepath.Join(storageDir, filename)
	return os.ReadFile(path)
}
