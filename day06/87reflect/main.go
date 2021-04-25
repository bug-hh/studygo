package main

import (
	"fmt"
	"reflect"
)

type Cat struct{}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", v)
	fmt.Printf("type name: %v, type kind: %v\n", v.Name(), v.Kind())
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本，reflect 包会引发 panic
	}
}

func reflectSetValue2(x interface{}) {
	// 由于 x 实际传入的时候，是指针，所以在用的时候，v.Elem()
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a float32 = 3.14
	var b int64 = 100
	var c = Cat{}

	reflectType(a)
	reflectType(b)
	reflectType(c)

	reflectSetValue2(&b)
	fmt.Println(b)

}
