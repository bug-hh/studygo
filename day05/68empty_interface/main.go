package main

import "fmt"

// interface{} 表示空接口
func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)

	m1["abc"] = "uiop"
	m1["aaa"] = 1234
	m1["bbb"] = true
	m1["ccc"] = [...]string{"唱", "跳", "rap"}

	fmt.Println(m1)

}
