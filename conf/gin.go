/**
 @author: RedCrazyGhost
 @date: 2023/4/9

**/

package conf

import "github.com/gin-gonic/gin"

var Gin *gin.Engine

func InitGin() {
	Gin = gin.New()
	Gin.Use(LogrusHandler(), gin.Recovery(), RBACHandler())

}
