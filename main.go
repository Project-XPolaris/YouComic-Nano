package main

import (
	"YouComic-Nano/config"
	"YouComic-Nano/datasource"
	"YouComic-Nano/debug"
	"YouComic-Nano/generate"
	"YouComic-Nano/router"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	var err error
	// debug mode
	debugMode := len(os.Getenv("DEBUG")) != 0
	rootPath := "./"
	if debugMode && len(os.Getenv(debug.DEBUG_GENERATE_PATH)) != 0 {
		rootPath = os.Getenv(debug.DEBUG_GENERATE_PATH)
	}
	if debugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}
	// log
	config.SetupLogConfig()
	config.RootPath = rootPath
	//clear up
	//err = generate.ClearUp()
	//if err != nil {
	//	logrus.Fatal(err)
	//}

	// load application config
	err = config.SetupApplicationConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	// load cache
	if config.ApplicationConfig.GetBool(config.ACTIVE_COVER_THUMBNAIL) {
		err = config.SetupCache()
		if err != nil {
			logrus.Fatal(err)
		}
	}

	// first library
	if generate.CheckIsNewLibrary(rootPath) {
		logrus.Info("build library ...")
		err := generate.SetupLibrary(rootPath)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	err = config.LoadConfig(rootPath)
	if err != nil {
		logrus.Fatal(err)
	}

	datasource.InitDataSource()
	r := gin.New()
	router.SetRouter(r)
	gin.SetMode(gin.DebugMode)
	err = r.Run("0.0.0.0:8880")
	if err != nil {
		logrus.Fatal(err)
	}
}
