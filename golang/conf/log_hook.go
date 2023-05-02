/**
@author: RedCrazyGhost
@date: 2023/4/5

**/

package conf

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

// Fire 自定义日志钩子 可能存在文件未关闭的情况，或是由程序自动GC回收关闭
func (hook LogHook) Fire(*logrus.Entry) error {
	fileName := time.Now().Format("2006-01-02")

	fullPath := fmt.Sprintf("%s/%s.log", Config.Log.DirPath, fileName)

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
