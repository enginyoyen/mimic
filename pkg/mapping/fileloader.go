package mapping

import (
	"errors"
	"os"
	"path/filepath"
)

func WalkFiles(rootPath string) (*[]string, error) {
	var files []string
	if len(rootPath) == 0 {
		return &files, errors.New("Empty folder name")
	}

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			files = append(files, path)
		}
		return nil
	})
	return &files, err
}
