package generate

import (
	"YouComic-Nano/config"
	"os"
	"path/filepath"
)

func ClearUp() error {
	err := os.RemoveAll(filepath.Join(config.RootPath, config.EXPORT_CONFIG_FILE_NAME))
	if err != nil {
		return err
	}
	err = os.RemoveAll(filepath.Join(config.RootPath, ".cache"))
	if err != nil {
		return err
	}
	return nil
}
