/**
 @author: RedCrazyGhost
 @date: 2023/4/16

**/

package controller

import (
	conf2 "EndProject/golang/conf"
	"EndProject/golang/core"
	"EndProject/golang/serivce"
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetData 废弃 数据遮盖
func GetData(c *gin.Context) {
	response := core.Response{C: c}
	var data []map[string]interface{}
	conf2.DB.Table("test_test").Select("col_0,col_1,col_2,col_3,col_4,col_5,col_6,concat(left(col_7,1),'*****',right(col_7,1)) as col_7,concat(left(col_8,1),'*****',right(col_8,1)) as col_8").Scan(&data)
	response.SuccessData(data)
}

// GetData1 废弃
func GetData1(c *gin.Context) {
	response := core.Response{C: c}
	var data []map[string]interface{}
	conf2.DB.Table("test_test").Select("date(col_4) as date,col_5 as price").Scan(&data)
	response.SuccessData(data)
}

func Upload(c *gin.Context) {
	response := core.Response{C: c}

	// 单文件
	file, _ := c.FormFile("file")
	userId, _ := c.GetQuery("UserId")

	dst := conf2.Config.Application.RootPath + "/file/" + userId + "_" + file.Filename
	// 上传文件至指定的完整文件路径
	err0 := c.SaveUploadedFile(file, dst)
	err1 := serivce.CreateDataTable(userId, file.Filename)
	if err0 != nil || err1 != nil {
		response.ErrorMsg(fmt.Sprintf("'%s' 失败原因：", file.Filename))
		return
	}
	response.SuccessMsg(fmt.Sprintf("'%s' 上传成功!", file.Filename))
}

func GetUserTables(c *gin.Context) {
	response := core.Response{C: c}
	userId, _ := c.GetQuery("UserId")
	tables := serivce.GetUserTables(userId)

	response.SuccessData(tables)
}

func CheckDataTable(c *gin.Context) {
	response := core.Response{C: c}
	TableName, ok := c.GetQuery("TableName")
	if !ok {
		response.ErrorMsg("缺少表名称")
		return
	}
	colBalance, ok := c.GetQuery("colBalance")
	if !ok {
		response.ErrorMsg("缺少余额列名")
		return
	}
	colAmount, ok := c.GetQuery("colAmount")
	if !ok {
		response.ErrorMsg("缺少交易金额列名")
		return
	}

	data := serivce.CheckDataTable(TableName, colBalance, colAmount)
	response.SuccessData(data)
}

func GetRowData(c *gin.Context) {
	response := core.Response{C: c}
	TableName, ok := c.GetQuery("TableName")
	if !ok {
		response.ErrorMsg("缺少表名称")
		return
	}
	rowNum, ok := c.GetQuery("RowNum")
	if !ok {
		response.ErrorMsg("缺少行数数据")
		return
	}
	data := serivce.GetRowData(TableName, rowNum)
	response.SuccessData(data)
}

func GetGroupCount(c *gin.Context) {
	response := core.Response{C: c}
	TableName, ok := c.GetQuery("TableName")
	if !ok {
		response.ErrorMsg("缺少表名称")
		return
	}
	groupCol, ok := c.GetQuery("GroupCol")
	if !ok {
		response.ErrorMsg("缺少行数数据")
		return
	}
	data := serivce.GetGroupCount(TableName, groupCol)
	response.SuccessData(data)

}

func GetTableCount(c *gin.Context) {
	response := core.Response{C: c}
	TableName, ok := c.GetQuery("TableName")
	if !ok {
		response.ErrorMsg("缺少表名称")
		return
	}
	data := serivce.GetTableCount(TableName)
	response.SuccessData(data)

}

func GetTableHead(c *gin.Context) {
	response := core.Response{C: c}
	TableName, ok := c.GetQuery("TableName")
	if !ok {
		response.ErrorMsg("缺少表名称")
		return
	}
	data := serivce.GetTableHead(TableName)
	response.SuccessData(data)

}
