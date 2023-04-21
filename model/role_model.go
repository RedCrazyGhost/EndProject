/**
 @author: RedCrazyGhost
 @date: 2023/4/6

**/

package model

type Role struct {
	Id     int    `gorm:"not null;autoIncrement;index;comment:'角色ID';primaryKey;"`
	Name   string `gorm:"unique;default:'';comment:'角色名'"`
	UserId uint64 `gorm:"type:text;comment:'用户ID'"`
}
