package main

import "fmt"

func funcA() {
	fmt.Println("A")
}

func funcB() {
	defer func() {
		// 调用 recover 函数，尝试从 panic 中恢复，那么panic后的语句不会执行，但是 funcB 函数后的代码会正常执行
		err := recover()
		fmt.Println(err)
		// 当出现严重错误后，可以在 defer 释放资源
		fmt.Println("释放数据库连接")
	}()
	panic("出现了严重的错误")
	fmt.Println("B")
}

func funcC() {
	fmt.Println("C")
}

func main() {
	funcA()
	funcB()
	funcC()
}
