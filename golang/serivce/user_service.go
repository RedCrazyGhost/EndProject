/**
 @author: RedCrazyGhost
 @date: 2023/4/10

**/

package serivce

import (
	"EndProject/golang/conf"
	model2 "EndProject/golang/model"
	"EndProject/golang/model/request"
	"errors"
	"time"
)

func Login(user *request.LoginUser) (*model2.User, error) {
	data := model2.User{}
	conf.DB.Table("users").Where(&user).Find(&data)
	if data.Email == "" {
		return nil, errors.New("用户不存在")
	}
	return &data, nil
}

func SaveUser(user model2.User) error {
	datetime := "2000-05-14 00:00:00"
	formatTime, _ := time.Parse("2006-01-02 15:04:05", datetime)
	user.LastLogin = formatTime
	conf.DB.Save(&user)
	conf.DB.Save(&model2.Role{UserId: user.ID, Name: "user"})
	return nil
}

func DeleteUser(user model2.User) error {
	user.IsDeleted = 1
	conf.DB.Save(&user)
	return nil
}
