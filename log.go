/**
 @author: RedCrazyGhost
 @date: 2023/4/5

**/

package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func InitLog() {
	rootPath, err := os.Getwd()
	if err != nil {
		Log.Panic("日志文件夹位置获取出错！", err)
	}
	if len(Config.Log.DirPath) == 0 || Config.Log.DirPath == "" {
		Config.Log.DirPath = rootPath + "/log"
	} else {
		Config.Log.DirPath = rootPath + "/" + Config.Log.DirPath
	}

	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetReportCaller(true)
	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	Log.AddHook(&LogHook{})
}
