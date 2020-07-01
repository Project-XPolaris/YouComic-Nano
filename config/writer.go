package config

import (
	"YouComic-Nano/entity"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

const (
	EXPORT_CONFIG_FILE_NAME = "export_library.json"
)

type LibraryExportConfig struct {
	Name  string          `json:"name"`
	Books []entity.Book `json:"books"`
}

func WriteExportLibraryConfig(books []entity.Book, name string, rootPath string) error {
	file, err := json.MarshalIndent(LibraryExportConfig{Books: books, Name: name}, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(rootPath, EXPORT_CONFIG_FILE_NAME), file, 0644)
	return err
}
