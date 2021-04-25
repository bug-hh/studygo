package main

import (
	"fmt"
	"os"
)

// 学生管理系统

var smr studentMgr

type student struct {
	id   int64
	name string
}

type studentMgr struct {
	allStudent map[int64]student
}

// 查看学生
func (s studentMgr) showStudent() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号: %d, 学生: %s\n", stu.id, stu.name)
	}
}

// 删除学生
func (s studentMgr) deleteStudent() {
	// 先获取用户输入学号
	var stuId int64
	fmt.Print("请输入要删除的学号: ")
	fmt.Scanln(&stuId)
	_, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	delete(s.allStudent, stuId)
	fmt.Println("删除成功")
}

// 增加学生
func (s studentMgr) addStudent() {
	// 根据用户输入创造一个学生
	var (
		stuId   int64
		stuName string
	)

	fmt.Print("请输入学号: ")
	fmt.Scanln(&stuId)
	fmt.Print("请输入学生姓名: ")
	fmt.Scanln(&stuName)
	newStu := student{
		id:   stuId,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
}

// 修改学生
func (s studentMgr) modifyStudent() {
	// 先获取用户输入学号
	var stuId int64
	fmt.Print("请输入学号: ")
	fmt.Scanln(&stuId)
	stuObj, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Print("请输入学生的新名字: ")
	var name string
	fmt.Scanln(&name)
	stuObj.name = name
	s.allStudent[stuId] = stuObj
}

func showMenu() {
	fmt.Println("Welcome to sms")
	fmt.Println(`
		1、查看所有学生
		2、增加学生
		3、删除学生
		4、修改学生
		5、退出
	`)
}

func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		fmt.Print("你的选择是：")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.deleteStudent()
		case 4:
			smr.modifyStudent()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("滚!")
		}
	}

}
