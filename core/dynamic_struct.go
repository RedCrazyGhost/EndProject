/**
 @author: RedCrazyGhost
 @date: 2023/4/6

**/

package core

import (
	"errors"
	"reflect"
)

// Struct
// map[string]int 用于通过字段名称从构造器的[]reflect.StructField中获取reflect.StructField
type Struct struct {
	Fields   []reflect.StructField
	Type     reflect.Type
	Instance reflect.Value
	Elem     reflect.Value
	index    map[string]int
}

// NewStruct 创建空Struct
func NewStruct() *Struct {
	return &Struct{}
}

// Build 根据预先添加的字段构建出结构体,并且对属性进行标记存储
// 并且生成实例
func (s *Struct) Build() error {
	if len(s.Fields) == 0 {
		return errors.New("请先通过AddField()方法生成属性！")
	}
	structOf := reflect.StructOf(s.Fields)
	s.index = make(map[string]int)
	for i := 0; i < structOf.NumField(); i++ {
		s.index[structOf.Field(i).Name] = i
	}
	s.Type = structOf

	if s.Type == nil {
		return errors.New("请先通过Addxxx()方法添加属性！")
	}
	s.Instance = reflect.New(s.Type)
	s.Elem = s.Instance.Elem()
	return nil
}

// GetFiled 获取属性值
func (i Struct) GetFiled(name string) (reflect.Value, error) {
	if index, ok := i.index[name]; ok {
		return i.Instance.Field(index), nil
	} else {
		return reflect.Value{}, errors.New("属性不存在！")
	}
}

// AddField 添加结构体字段信息
func (s *Struct) AddField(fieldName string, Type reflect.Type, tag string) {
	s.Fields = append(s.Fields, reflect.StructField{
		Name: fieldName,
		Type: Type,
		Tag:  reflect.StructTag(tag),
	})
}

// AddString 添加string类型
func (s *Struct) AddString(name string, tag string) {
	s.AddField(name, reflect.TypeOf(""), tag)
}

// AddInt64 添加Int64类型
func (s *Struct) AddInt64(name string, tag string) {
	s.AddField(name, reflect.TypeOf(int64(0)), tag)
}

// InterfaceOfValue 获取interface
func (s *Struct) InterfaceOfValue() interface{} {
	return s.Instance.Interface()
}

// AddrOfValue 获取地址值
func (s *Struct) AddrOfValue() interface{} {
	return s.Instance.Addr()
}

// ElemOfValue 获取Elem
func (s *Struct) ElemOfValue() interface{} {
	return s.Elem
}

// GetTag 获取属性标签
func (i Struct) GetTag(name, tagName string) string {
	if field, ok := i.Type.FieldByName(name); ok {
		return field.Tag.Get(tagName)
	}
	return ""
}

// SetString 设置Instance值
func (s *Struct) SetString(name, value string) {
	if index, ok := s.index[name]; ok {
		s.Elem.Field(index).SetString(value)
	}
}

// SetInt64 设置Instance值
func (s *Struct) SetInt64(name string, value int64) {
	if index, ok := s.index[name]; ok {
		s.Elem.Field(index).SetInt(value)
	}
}
