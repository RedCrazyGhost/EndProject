/**
 @author: RedCrazyGhost
 @date: 2023/4/10

**/

package request

type LoginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
