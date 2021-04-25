package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 100,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())
	// 通过 for 循环遍历结构体所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name: %s, index: %d, type: %v, json tag: %v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名，获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name: %s, index: %d, type: %v, json tag: %v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
