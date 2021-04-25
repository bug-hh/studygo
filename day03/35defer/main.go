package main

import "fmt"

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println(1) // defer 语句把它作用的语句延迟到函数即将返回时逆序执行
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

func main() {
	fmt.Println(f1())
}

// 在 go 中，函数返回操作不是原子操作，分成两步：
// 第一步：给返回值赋值
// 第二步：返回返回值
// 那么 defer 语句在第一步与第二步之间执行
func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是 x，不是返回值
	}()
	return x
}

// 请问 f1 的返回值是多少?  5

func f2() (x int) {
	x = 5
	defer func() {
		x++
	}()
	return x
}

// 请问 f2 的返回值是多少?  6
