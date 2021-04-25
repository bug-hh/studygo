package main

import "fmt"

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1)
	i2 := 077
	fmt.Printf("%d\n", i2) // 八进制转十进制输出
	i3 := 0x1234
	fmt.Printf("%x\n", i3) // 打印十六进制
	// 使用 %T 查看变量类型
	fmt.Printf("%T\n", i3)
}
