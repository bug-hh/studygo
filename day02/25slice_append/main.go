package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := a[:] // b 是切片
	// 删除数字 3
	b = append(b[:2], b[3:]...)
	fmt.Println(b) // 1，2，4，5，6，7，8
	fmt.Println(a) // 1, 2, 4, 5,6,7,8,8
}
