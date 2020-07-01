package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
)

var CacheConfig *viper.Viper

func SetupCache() error {
	cachePath := filepath.Join(RootPath, ".cache")
	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		err = InitCache()
		if err != nil {
			return err
		}
	}
	CacheConfig = viper.New()
	CacheConfig.SetConfigName("manifest")
	CacheConfig.SetConfigType("json")
	CacheConfig.AddConfigPath(filepath.Join(RootPath, ".cache"))
	err := CacheConfig.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func InitCache() error {
	cachePath := filepath.Join(RootPath, ".cache")
	err := os.MkdirAll(cachePath, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := json.MarshalIndent(map[string]string{}, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(cachePath, "manifest.json"), file, 0644)
	return err
}
