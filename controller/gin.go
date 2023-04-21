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
	Gin.Use(conf.LogrusHandler(), gin.Recovery(), conf.RBACHandler())

	Gin.POST("/login", Login)

	//用户
	{
		//Gin.GET()
	}

	//表数据
	data := Gin.Group("/data")
	{
		data.GET("/getdata", GetData)
		data.GET("/getdata1", GetData1)
		data.GET("/getdata2", GetData2)
		data.POST("/upload", Upload)
		data.GET("/GetUserTables", GetUserTables)
	}

	err := Gin.Run()
	if err != nil {
		conf.Log.Panicf("服务器启动失败！失败原因：%v", err)
	}

}
