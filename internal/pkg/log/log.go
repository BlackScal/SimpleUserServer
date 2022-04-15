package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
	appid  string
)

type LogFields map[string]interface{}

func SetFormatter(format string) {
	if format == "json" {
		logger.Formatter = &logrus.JSONFormatter{}
	}
}

func SetOutput(output string) {
	if output == "console" {
		logger.Out = os.Stdout
	} else { //TODO
		_ = output
	}
}

func SetLevel(levelstr string) {
	var level logrus.Level

	switch levelstr {
	case "debug":
		level = logrus.WarnLevel
	case "info":
		level = logrus.InfoLevel
	case "warning":
		level = logrus.WarnLevel
	default:
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
}

func SetAppid(aid string) {
	appid = aid
}

func Debug(kv LogFields, msg string) {
	kv["appid"] = appid
	log(logrus.DebugLevel, kv, msg)
}

func Info(kv LogFields, msg string) {
	kv["appid"] = appid
	log(logrus.InfoLevel, kv, msg)
}

func Warn(kv LogFields, msg string) {
	kv["appid"] = appid
	log(logrus.WarnLevel, kv, msg)
}

func Error(kv LogFields, msg string) {
	kv["appid"] = appid
	log(logrus.ErrorLevel, kv, msg)
}

func log(lvl logrus.Level, kv LogFields, msg string) {
	kv["appid"] = appid
	logger.WithFields(logrus.Fields(kv)).Log(lvl, msg)
}
