/**
 @author: RedCrazyGhost
 @date: 2023/4/6

**/

package model

import (
	"EndProject/golang/conf"
	"time"
)

type BaseModel struct {
	ID        uint64    `gorm:"not null;autoIncrement;index;comment:'唯一ID';primaryKey;"`
	CreatedBy uint64    `gorm:"comment:'创建人';"`
	CreatedAt time.Time `gorm:"comment:'创建时间';"`
	UpdatedBy uint64    `gorm:"comment:'修改人';"`
	UpdatedAt time.Time `gorm:"comment:'修改时间';"`
	Remark    string    `gorm:"type:varchar(255);comment:'备注';"`
	IsDeleted int8      `gorm:"default:0;not null;index;comment:'逻辑删除 0-未删除 1-已删除'"`
}

func InitModel() {
	conf.DB.Table("users").AutoMigrate(&User{})
	conf.CreateTable("users", &User{})
	conf.DB.Table("roles").AutoMigrate(&Role{})
	conf.CreateTable("roles", &Role{})
}
