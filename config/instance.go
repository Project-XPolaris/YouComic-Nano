package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Current LibraryExportConfig
var RootPath string
func LoadConfig(rootPath string) error {
	configPath := filepath.Join(rootPath, EXPORT_CONFIG_FILE_NAME)
	configFile, err := os.Open(configPath)
	if err != nil {
		return err
	}
	configFileStream, err := ioutil.ReadAll(configFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(configFileStream, &Current)
	return err
}
