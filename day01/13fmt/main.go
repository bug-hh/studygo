package main

import "fmt"

func main() {
	a := 100
	b := 1.234
	fmt.Printf("%d\n", a) // 十进制
	fmt.Printf("%o\n", a) // 八进制
	fmt.Printf("%x\n", a) // 十六进制
	fmt.Printf("%v\n", a) // 打印出变量的值
	fmt.Printf("%b\n", a) // 二进制
	fmt.Printf("%T\n", a) // 打印出变量的类型
	c := "abc"
	fmt.Printf("%f\n", b) // 浮点数
	fmt.Printf("%s\n", c) // 字符串
	fmt.Printf("%#v\n", c)
}
