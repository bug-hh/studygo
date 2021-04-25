package main

import "fmt"

type person struct {
	name string
	age  int
}

// 构造函数
// 返回结构体还是结构体指针？
// 当结构体比较大的时候，尽量使用结构体指针，减少程序运行开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}

}
func main() {
	// 匿名结构体
	var s struct {
		x string
		y int
	}

	s.x = "呵呵"
	s.y = 10

	fmt.Printf("%T, %v\n", s, s)

	// 使用 & 实例化结构体
	var n = &person{
		name: "phh",
		age:  27,
	}

	fmt.Printf("%v\n", n)

	var m = newPerson("bughh", 27)
	fmt.Printf("%v\n", m)
}
