package main

import "fmt"

type myInt int     // 自定义类型
type yourInt = int // 类型别名

func main() {
	var n myInt
	n = 10
	var m yourInt
	m = 30
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", m)
}
