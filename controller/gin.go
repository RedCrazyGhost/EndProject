/**
 @author: RedCrazyGhost
 @date: 2023/4/9

**/

package controller

import (
	"EndProject/conf"
	"github.com/gin-gonic/gin"
)

var Gin *gin.Engine

func InitGin() {
	Gin = gin.New()
	Gin.Use(conf.LogrusHandler(), gin.Recovery(), conf.CorsHandler(), conf.RBACHandler())

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
		conf.Log.Panicf("服务器启动失败！失败原因：%v", err)
	}

}
