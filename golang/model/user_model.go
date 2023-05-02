/**
 @author: RedCrazyGhost
 @date: 2023/4/6

**/

package model

import "time"

type User struct {
	BaseModel `gorm:"embedded;"`
	UserName  string    `gorm:"type:varchar(40);size 40;not null;comment:'用户名'"`
	Password  string    `gorm:"type:varchar(200);not null;size 200;comment:'密码'"`
	RealName  string    `gorm:"type:varchar(50);size 50;comment:'真实姓名'"`
	Email     string    `gorm:"type:varchar(50);size 50;not null;comment:'邮箱'"`
	LastLogin time.Time `gorm:"comment:'最近一次登录'"`
	Phone     string    `gorm:"type:varchar(11);size 11;comment:'电话'"`
	IdCard    string    `gorm:"type:varchar(18);size 18;comment:'身份证号'"`
	Gender    int8      `gorm:"default:1;comment:'性别 女-0 男-1'"`
	Status    int8      `gorm:"default:1;comment:'状态 正常-1 封禁-2'"`
}
