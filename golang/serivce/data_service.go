/**
 @author: RedCrazyGhost
 @date: 2023/4/20

**/

package serivce

import (
	conf2 "EndProject/golang/conf"
	"EndProject/golang/core"
	"fmt"
	"github.com/xuri/excelize/v2"
	"regexp"
	"strings"
)

func GetTableCount(table string) map[string]interface{} {
	var data map[string]interface{}
	sql := `
select count(1) as count from #{table};
`
	sql = strings.ReplaceAll(sql, "#{table}", table)
	conf2.DB.Raw(sql).Scan(&data)
	return data
}

//select  month(col_4),max(#{col_balance})
//from test_test
//group by month(col_4)
// 每月最大余额

func GetTableStruct(table string) {
	var tableStruct map[string]interface{}
	conf2.DB.Raw("show create table ?", table).Scan(tableStruct)
	tabledata := tableStruct["Create Table"]
	reg := regexp.MustCompile("(?<=`).*(?=`)")
	if reg == nil { //解释失败，返回nil
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg.FindAllStringSubmatch(tabledata.(string), -1)
	fmt.Println("result1 = ", result1)
}

// CreateDataTable 创建数据表并且填充数据
func CreateDataTable(userId, filename string) error {
	f, err := excelize.OpenFile(conf2.Config.Application.RootPath + "/file/" + userId + "_" + filename)
	if err != nil {
		return err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//Get all the rows in the Sheet1.
	SheetName := "Sheet1"
	rows, err := f.GetRows(SheetName)
	if err != nil {
		return err
	}

	//table表名规则 UserId_ExelTableName_sheetName
	newStruct := conf2.TableStruct{Struct: core.NewStruct(), TableName: fmt.Sprintf("%s_%s_%s", userId, strings.Split(filename, ".")[0], SheetName)}
	dataMap := make([]map[string]interface{}, 0)
	for rowIndex, row := range rows {
		ColMap := make(map[string]interface{}, 0)
		for coli, colCell := range row {
			if rowIndex == 0 {
				newStruct.AddString(
					fmt.Sprintf("Col_%d", coli),
					fmt.Sprintf(`gorm:"type:varchar(255);size 255;comment:%s"`, colCell),
				)
			} else {
				if coli >= len(newStruct.Fields) {
					newStruct.AddString(
						fmt.Sprintf("Col_%d", len(newStruct.Fields)),
						fmt.Sprintf(`gorm:"type:varchar(255);size 255;comment:无名列-%d"`, len(newStruct.Fields)-coli),
					)
					err := newStruct.Build()
					if err != nil {
						return err
					}
				}
				ColMap[newStruct.Fields[coli].Name] = colCell
			}
		}
		if rowIndex == 0 {
			err := newStruct.Build()
			if err != nil {
				return err
			}
		} else {
			dataMap = append(dataMap, ColMap)
		}
	}

	newStruct.CreateTable()
	var count int64
	newStruct.Table.Count(&count)
	if count == 0 {
		newStruct.Table.Create(dataMap)
	}
	return nil
}

// GetUserTables 获取用户所拥有的表
func GetUserTables(userId string) []map[string]interface{} {
	var data []map[string]interface{}
	sql := `SELECT distinct TABLE_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME LIKE ?;`
	conf2.DB.Raw(sql, userId+"_%").Scan(&data)
	return data
}

// CheckDataTable 检测表中计算错误
func CheckDataTable(tableName, colBalance, colAmount string) []map[string]interface{} {
	sql := `
select ROUND( ROUND(cast(t1.#{col_balance} as float ),2)-ROUND(cast(t1.#{col_amount} as float ),2),2) as checkResult,ROUND(cast(t2.#{col_balance} as float ),2) as failResult,
       t1.*
from (select(@rowNum1:=@rowNum1+1) as rowNum,t.* from #{table} t,(select @rowNum1:=0) as t0) t1
left join (select (@rowNum2:=@rowNum2+1) as rowNum,t.* from #{table} t,(select @rowNum2:=0) as t0) t2
on t1.rowNum=t2.rowNum+1
where ROUND( ROUND(cast(t1.#{col_balance} as float ),2)-ROUND(cast(t1.#{col_amount} as float ),2),2)!=ROUND(cast(t2.#{col_balance} as float ),2);
`
	sql = strings.ReplaceAll(sql, "#{table}", tableName)
	//col_balance 余额
	sql = strings.ReplaceAll(sql, "#{col_balance}", colBalance)
	//col_amount 交易金额
	sql = strings.ReplaceAll(sql, "#{col_amount}", colAmount)

	var data []map[string]interface{}
	conf2.DB.Raw(sql).Scan(&data)
	return data
}

// GetRowData 获取表中某行数据
func GetRowData(tableName, rowNum string) map[string]interface{} {
	sql := `
select * from (select(@rowNum1:=@rowNum1+1) as rowNum,t.* from #{table} t,(select @rowNum1:=0) as t0) as t1 where rowNum=#{rowNum};`
	sql = strings.ReplaceAll(sql, "#{table}", tableName)
	sql = strings.ReplaceAll(sql, "#{rowNum}", rowNum)
	var data map[string]interface{}
	conf2.DB.Raw(sql).Scan(&data)
	return data
}

// GetGroupCount 分组聚合数量
func GetGroupCount(tableName, groupCol string) []map[string]interface{} {
	sql := `
select #{groupCol} as k ,count(#{groupCol}) as v from #{table} group by #{groupCol};
`
	sql = strings.ReplaceAll(sql, "#{table}", tableName)
	sql = strings.ReplaceAll(sql, "#{groupCol}", groupCol)
	var data []map[string]interface{}
	conf2.DB.Raw(sql).Scan(&data)
	return data
}

func GetTableHead(tableName string) []map[string]interface{} {
	sql := `select COLUMN_NAME as col_name, COLUMN_COMMENT col_comment from INFORMATION_SCHEMA.COLUMNS Where table_name ='#{table}';`
	sql = strings.ReplaceAll(sql, "#{table}", tableName)
	var data []map[string]interface{}
	conf2.DB.Raw(sql).Scan(&data)
	return data
}
