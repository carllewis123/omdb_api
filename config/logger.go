package config

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	filename "github.com/keepeye/logrus-filename"
	logrus "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

var logs = logrus.New()

type ConfigLog struct {
	Directory string `yaml:"log_file_dir"`
	Filename  string `yaml:"log_file_name"`
}

func ConfigLogger() (conf *ConfigLog, err error) {
	data, err := ioutil.ReadFile("./config/log_config.yaml")
	if err != nil {
		return
	}
	conf = &ConfigLog{}
	err = yaml.Unmarshal(data, conf)
	return
}

func Logger() *logrus.Logger {
	var configLog, _ = ConfigLogger()

	t := time.Now()
	t.String()
	date := t.Format("2006_01_02")
	var logLevel = logrus.DebugLevel
	var logs = logrus.New()
	var LogFilenName = configLog.Directory + "" + configLog.Filename + "_" + date + ".txt"
	f, err := os.OpenFile(LogFilenName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	mw := io.MultiWriter(os.Stdout, f)

	filenameHook := filename.NewHook()
	filenameHook.Field = "line"
	logs.AddHook(filenameHook)

	logs.SetFormatter(&logrus.TextFormatter{})
	logs.SetOutput(mw)
	logs.SetLevel(logLevel)

	return logs
}
