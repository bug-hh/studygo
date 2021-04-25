package main

import "fmt"

func f2() int {
	return 6
}

// 函数作为参数传递
func f1(x func() int) {
	a := x()
	fmt.Println(a)
}

func main() {
	f1(f2)
}
