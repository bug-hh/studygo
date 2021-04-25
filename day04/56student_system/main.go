package main

import (
	"fmt"
	"os"
)

type student struct {
	name string
	id   int
}

func main() {
	for {
		fmt.Println("欢迎来到学生管理系统")
		fmt.Println(`
		1、查看所有学生
		2、新增学生
		3、删除学生
		4、退出
	`)
		fmt.Print("请输入你想干嘛：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了 %d 这个选项", choice)
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("滚")
		}
	}

}

func showAllStudent() {
	fmt.Println("show")
}

func addStudent() {
	fmt.Println("add")
}

func deleteStudent() {
	fmt.Println("delete")
}
