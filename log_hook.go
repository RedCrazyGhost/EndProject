/**
@author: RedCrazyGhost
@date: 2023/4/5

**/

package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

type LogHook struct {
}

func (hook LogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

var LogFile *os.File

func (hook LogHook) Fire(*logrus.Entry) error {
	fileName := time.Now().Format("2006-01-02")

	fullPath := fmt.Sprintf("%s/%s.log", Config.Log.DirPath, fileName)

	// 无需多次获取文件句柄
	if LogFile != nil && LogFile.Name() == fullPath {
		return nil
	}

	if err := os.MkdirAll(Config.Log.DirPath, os.ModePerm); err != nil {
		Log.Panic("创建文件夹错误！", err)
		return err
	}

	LogFile, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Panic("写入日志文件错误！", err)
		return err
	}

	Log.Out = io.MultiWriter([]io.Writer{LogFile, os.Stdout}...)
	return nil
}
