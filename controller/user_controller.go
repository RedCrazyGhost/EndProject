/**
 @author: RedCrazyGhost
 @date: 2023/4/6

**/

package controller

import (
	"EndProject/core"
	"EndProject/model"
	"EndProject/serivce"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SaveUser(c *gin.Context) {
	response := core.Response{C: c}
	user := model.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.ErrorMsg(err.Error())
		return
	}
	err := serivce.SaveUser(user)
	if err != nil {
		response.ErrorMsg(fmt.Sprintf("创建/更新用户失败！%v", err))
		return
	}

	response.SuccessMsg("创建/更新账号成功！")
}
