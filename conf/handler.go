/**
 @author: RedCrazyGhost
 @date: 2023/4/9

**/

package conf

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//sub := "user"   // 想要访问资源的用户。
//	obj := "/admin" // 将被访问的资源。
//	act := "GET"    // 用户对资源执行的操作。
//	ok, err := conf.RBAC.Enforce(sub, obj, act)
//	if err != nil {
//		return
//	}
//	if ok == true {
//		fmt.Print("YES!")
//	} else {
//		fmt.Println("NO!")
//	}

func RBACHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := ""
		method := c.Request.Method
		path := c.Request.URL.Path

		ok, err := RBAC.Enforce(role, path, method)
		if err != nil {
			Log.Infof("RBAC执行发生错误！错误原因：%v", err)
		}

		if !ok {
			c.JSON(http.StatusUnauthorized, "权限不足！")
		}
	}
}
