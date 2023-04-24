/**
 @author: RedCrazyGhost
 @date: 2023/4/10

**/

package controller

import (
	"EndProject/core"
	"EndProject/model/request"
	"EndProject/serivce"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	response := core.Response{C: c}
	user := request.LoginUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ErrorMsg(err.Error())
		return
	}

	data, err := serivce.Login(&user)
	if err != nil {
		response.ErrorMsg(err.Error())
		return
	}
	response.SuccessData(data)
}
