package main

import "fmt"

var x = 100

func f1() {
	// 先找局部变量
	// 再找全局变量
	fmt.Println(x)
}
func main() {
	f1()

}
