package generate

import (
	"YouComic-Nano/config"
	"os"
	"path/filepath"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func CheckIsNewLibrary(libraryPath string) bool {
	return !fileExists(filepath.Join(libraryPath, config.EXPORT_CONFIG_FILE_NAME))
}
func SetupLibrary(libraryPath string) error {
	books, err := CreateLibrary(libraryPath)
	if err != nil {
		return err
	}
	err = config.WriteExportLibraryConfig(books, filepath.Base(libraryPath), libraryPath)
	if err != nil {
		return err
	}

	err = GenerateLibraryCoverThumbnail(books)
	if err != nil {
		return err
	}
	return err
}
