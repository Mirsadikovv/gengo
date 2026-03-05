package sharedutil

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
)

type JsonArray []map[string]any

func (j JsonArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JsonArray) Scan(value any) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("unsupported data type: %T", value)
	}
	return json.Unmarshal(bytes, j)
}

type JsonObject map[string]any

func (j JsonObject) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JsonObject) Scan(value any) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("unsupported data type: %T", value)
	}
	return json.Unmarshal(bytes, j)
}

func MustValue[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func WalkDirWithWriter(fsys fs.FS, dir, ext string, writer io.Writer) error {

	return fs.WalkDir(fsys, dir, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if d.IsDir() || filepath.Ext(path) != ext {
			return nil
		}

		f, openErr := fsys.Open(path)
		{
			if openErr != nil {
				return openErr
			}
			defer f.Close()
		}

		if _, err := io.Copy(writer, f); err != nil {
			return err
		}

		return nil
	})
}

func WalkDir(fsys fs.FS, dir, ext string) (result []string, err error) {

	fs.WalkDir(fsys, dir, func(path string, d fs.DirEntry, direrr error) error {

		if direrr != nil {
			err = errors.Join(err, direrr)
			return direrr
		}

		if d.IsDir() || filepath.Ext(path) != ext {
			return nil
		}

		f, openErr := fsys.Open(path)
		{
			if openErr != nil {
				err = errors.Join(err, openErr)
				return openErr
			}
			defer f.Close()
		}

		data, direrr := io.ReadAll(f)
		{
			if direrr != nil {
				err = errors.Join(err, direrr)
				return direrr
			}
		}

		result = append(result, string(data))

		return nil
	})

	return
}
