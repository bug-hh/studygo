package main

import "fmt"

type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	// 匿名嵌套结构体
	address
}

func main() {
	p1 := person{
		name: "abc",
		age:  18,
		address: address{
			province: "北京",
			city:     "北京",
		},
	}
	// 由于 address 是一个匿名嵌套结构体，p1 可以直接将 address 的成员当做自己的成员来访问
	fmt.Println(p1.province)
}
