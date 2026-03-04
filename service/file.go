package service

import (
	"bytes"
	"os"
	"path/filepath"
)

// RenderToFile renders a template string with data and writes it to a file path.
func RenderToFile(tmplStr string, data map[string]interface{}, parts ...string) error {
	tmpl := Tmp(tmplStr)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	filePath := filepath.Join(parts...)
	return os.WriteFile(filePath, buf.Bytes(), 0644)
}
