package main

import "fmt"

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s 会动\n", a.name)
}

type dog struct {
	feet int
	animal
}

func (d dog) wang() {
	// 由于 dog 结构体中没有 name 成员，于是去嵌套结构体 animal 中找
	fmt.Printf("%s 在叫，汪汪汪\n", d.name)
}

func main() {
	d := dog{
		feet: 4,
		animal: animal{
			name: "abc",
		},
	}
	// 由于 dog 结构体中没有 move 方法，于是去嵌套结构体 animal 中找
	d.move()
	d.wang()
}
