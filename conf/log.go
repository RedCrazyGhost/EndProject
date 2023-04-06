/**
 @author: RedCrazyGhost
 @date: 2023/4/5

**/

package conf

import (
	"github.com/sirupsen/logrus"
	"os"
)

type LogConfig struct {
	DirPath string
}

var Log *logrus.Logger

func InitLog() {
	if len(Config.Log.DirPath) == 0 || Config.Log.DirPath == "" {
		Config.Log.DirPath = Config.Application.RootPath + "/log"
	} else {
		Config.Log.DirPath = Config.Application.RootPath + "/" + Config.Log.DirPath
	}

	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetReportCaller(true)
	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	Log.AddHook(&LogHook{})
}
