/**
 @author: RedCrazyGhost
 @date: 2023/4/9

**/

package conf

import (
	"EndProject/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RBACHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var roleData map[string]interface{}
		var role string

		UserId, _ := c.GetQuery("UserId")
		if len(UserId) != 0 {
			DB.Raw("select * from roles where user_id = ?", UserId).Scan(&roleData)
			if roleData == nil {
				role = ""
			} else {
				role = roleData["name"].(string)
			}
		} else {
			role = ""
		}

		method := c.Request.Method
		path := c.Request.URL.Path

		ok, err := RBAC.Enforce(role, path, method)
		if err != nil {
			Log.Infof("RBAC执行发生错误！错误原因：%v", err)
		}

		if !ok {
			response := core.Response{C: c}
			response.Send(http.StatusUnauthorized, "权限不足！", nil)
			c.Abort()
		}
	}
}
