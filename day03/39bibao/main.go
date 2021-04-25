package main

import "fmt"

func f1(x int) func(int) int {
	// return 的这个函数，就是一个闭包，闭包是一个函数，这个函数包含了它外部作用域的一个变量（这里是 x）
	return func(z int) int {
		x += z
		return x
	}
}

func main() {
	ret := f1(100)
	ret2 := ret(500)
	fmt.Println(ret2)
}
