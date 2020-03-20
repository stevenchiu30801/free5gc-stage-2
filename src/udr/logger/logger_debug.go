//+build debug

package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gofree5gc/lib/logger_util"
	"gofree5gc/lib/path_util"
	"os"
	"runtime"
	"strings"
)

var log *logrus.Logger
var AppLog *logrus.Entry
var InitLog *logrus.Entry
var HandlerLog *logrus.Entry
var DataRepoLog *logrus.Entry

func init() {
	log = logrus.New()
	log.SetReportCaller(true)

	log.Formatter = &logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			orgFilename, _ := os.Getwd()
			repopath := orgFilename
			repopath = strings.Replace(repopath, "/bin", "", 1)
			filename := strings.Replace(f.File, repopath, "", -1)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	fileHook, err := logger_util.NewFileHook(path_util.Gofree5gcPath("gofree5gc/free5gc.log"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		log.Hooks.Add(fileHook)
	}

	AppLog = log.WithFields(logrus.Fields{"UDR": "app"})
	InitLog = log.WithFields(logrus.Fields{"UDR": "init"})
	HandlerLog = log.WithFields(logrus.Fields{"UDR": "Handler"})
	DataRepoLog = log.WithFields(logrus.Fields{"UDR": "DataRepo"})
}

func SetLogLevel(level logrus.Level) {
	log.SetLevel(level)
}

func SetReportCaller(bool bool) {
	log.SetReportCaller(bool)
}
