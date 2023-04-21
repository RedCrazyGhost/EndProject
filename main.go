/**
 @author: RedCrazyGhost
 @date: 2023/4/5

**/

package main

import (
	"EndProject/conf"
	"EndProject/controller"
	"EndProject/model"
)

func init() {
	conf.InitConfig()
	conf.InitLog()
	conf.InitCasbin()
	conf.InitDB()
	model.InitModel()
	controller.InitGin()
}

func main() {

}
