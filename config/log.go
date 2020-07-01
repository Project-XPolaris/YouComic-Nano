package config

import (
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

func SetupLogConfig() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())
}