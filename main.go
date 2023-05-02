/**
 @author: RedCrazyGhost
 @date: 2023/4/5

**/

package main

import (
	conf2 "EndProject/golang/conf"
	"EndProject/golang/controller"
	"EndProject/golang/model"
)

func init() {
	conf2.InitConfig()
	conf2.InitLog()
	conf2.InitCasbin()
	conf2.InitDB()
	model.InitModel()
	controller.InitGin()
}

func main() {

}
