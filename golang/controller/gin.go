/**
 @author: RedCrazyGhost
 @date: 2023/4/9

**/

package controller

import (
	conf2 "EndProject/golang/conf"
	"github.com/gin-gonic/gin"
)

var Gin *gin.Engine

func InitGin() {
	Gin = gin.New()
	Gin.Use(conf2.LogrusHandler(), gin.Recovery(), conf2.CorsHandler(), conf2.RBACHandler())

	Gin.POST("/login", Login)

	//用户
	user := Gin.Group("/user")
	{
		user.POST("/save", SaveUser)
	}

	//表数据
	data := Gin.Group("/data")
	{
		data.POST("/upload", Upload)
		data.GET("/list", GetUserTables)
		data.GET("/check", CheckDataTable)
		data.GET("/row", GetRowData)
		data.GET("/group", GetGroupCount)
		data.GET("/count", GetTableCount)
		data.GET("/head", GetTableHead)
	}

	err := Gin.Run()
	if err != nil {
		conf2.Log.Panicf("服务器启动失败！失败原因：%v", err)
	}

}
