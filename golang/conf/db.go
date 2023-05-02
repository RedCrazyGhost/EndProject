/**
@author: RedCrazyGhost
@date: 2023/4/5

**/

package conf

import (
	"EndProject/golang/core"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DSN string
}

var DB *gorm.DB

func InitDB() {
	var err error
	Log.Info("数据库连接开始！")
	DB, err = gorm.Open(mysql.Open(Config.Database.DSN), &gorm.Config{
		Logger: New(),
	})
	if err != nil {
		Log.Panicf("数据库连接失败！失败原因：%v 可能存在问题，请查看./config.yml文件!", err)
	}
	Log.Info("数据库连接成功！")
}

type TableStruct struct {
	*core.Struct
	TableName string
	Table     *gorm.DB
	Comment   string
}

// AddTable 创建数据库表
func (t *TableStruct) AddTable() {
	t.Table = DB.Table(t.TableName).Set(
		"gorm:table_options",
		fmt.Sprintf("ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '%s'", t.Comment))
}

// CreateTable 创建数据库表
func (t *TableStruct) CreateTable() {
	t.Table = DB.Table(t.TableName).Set(
		"gorm:table_options",
		fmt.Sprintf("ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '%s'", t.Comment))
	CreateTable(t.TableName, t.InterfaceOfValue())
}

// CreateTable 根据表名和struct创建或迁移表
func CreateTable(name string, s interface{}) {
	if !DB.Migrator().HasTable(s) {
		Log.Infof("创建%s表格开始！", name)
		err := DB.Table(name).AutoMigrate(s)
		if err != nil {
			Log.Panicf("创建%s表是失败！失败原因为：%v", name, err)
		}
		Log.Infof("创建%s表格成功！", name)
	}
}
