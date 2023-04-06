/**
 @author: RedCrazyGhost
 @date: 2023/4/5

**/

package conf

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type SystemConfig struct {
	Application ApplicationConfig
	Server      ServerConfig
	Database    DatabaseConfig
	Log         LogConfig
	Casbin      CasbinConfig
}
type ApplicationConfig struct {
	Name     string
	RootPath string
}
type ServerConfig struct {
	Port string
}

var Config *SystemConfig

func InitConfig() {
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

	if len(Config.Application.RootPath) == 0 || Config.Application.RootPath == "" {
		rootPath, err := os.Getwd()
		if err != nil {
			Log.Panicf("应用根地址获取失败！失败原因：%v", err)
		}
		Config.Application.RootPath = rootPath
	}

	Log.Infof("配置初始化完成！配置数据为：%v", Config)

}
