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

// CorsHandler  跨域配置
func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

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
