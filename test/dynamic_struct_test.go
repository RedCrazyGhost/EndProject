/**
 @author: RedCrazyGhost
 @date: 2023/4/6

**/

package test

import (
	"EndProject/core"
	"fmt"
	"testing"
)

func TestUse(t *testing.T) {
	newStruct := core.NewStruct()
	newStruct.AddString("Name", `gorm:"not null"`)
	newStruct.AddInt64("ID", `gorm:"primaryKey;"`)
	err := newStruct.Build()
	if err != nil {
		return
	}
	newStruct.SetString("Name", "RedCrazyGhost")
	newStruct.SetInt64("ID", 1)

	fmt.Print(newStruct.Elem)

}
