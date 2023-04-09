/**
 @author: RedCrazyGhost
 @date: 2023/4/6

**/

package conf

import "github.com/casbin/casbin/v2"

type CasbinConfig struct {
	ModelPath  string
	PolicyPath string
}

var RBAC *casbin.Enforcer

func InitCasbin() {
	if len(Config.Casbin.ModelPath) == 0 || Config.Casbin.ModelPath == "" {
		Config.Casbin.ModelPath = Config.Application.RootPath + "/model.conf"
	} else {
		Config.Casbin.ModelPath = Config.Application.RootPath + "/" + Config.Casbin.ModelPath
	}

	if len(Config.Casbin.PolicyPath) == 0 || Config.Casbin.PolicyPath == "" {
		Config.Casbin.PolicyPath = Config.Application.RootPath + "/policy.csv"
	} else {
		Config.Casbin.PolicyPath = Config.Application.RootPath + "/" + Config.Casbin.PolicyPath
	}

	var err error
	RBAC, err = casbin.NewEnforcer(Config.Casbin.ModelPath, Config.Casbin.PolicyPath)
	if err != nil {
		Log.Panicf("角色配置权限失败！失败原因：%v", err)
	}
}
