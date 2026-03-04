package service

import (
	"os"
	"path/filepath"
)

func MkdirAll(parts ...string) error {
	return os.MkdirAll(filepath.Join(parts...), os.ModePerm)
}
