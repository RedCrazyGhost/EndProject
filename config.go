/**
 @author: RedCrazyGhost
 @date: 2023/4/5

**/

package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type SystemConfig struct {
	Application ApplicationConfig
	Server      ServerConfig
	Datebase    DatabaseConfig
	Log         LogConfig
}
type ApplicationConfig struct {
	Name string
}
type ServerConfig struct {
	Port string
}
type DatabaseConfig struct {
	DSN string
}
type LogConfig struct {
	DirPath string
}

var Config *SystemConfig

func initConfig() {
	Log = logrus.New()

	Log.Info("配置初始化开始！")
	v := viper.New()
	v.SetConfigFile("./config.yml")
	err := v.ReadInConfig()
	if err != nil {
		Log.Panicf("配置文件读取失败！失败原因：%v，请正确配置文件./config.yml！", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		Log.Panicf("配置文件解析失败！失败原因：%v", err)
	}
	Log.Infof("配置初始化完成！配置数据为：%v", Config)

}
