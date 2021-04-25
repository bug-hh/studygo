package main

import "fmt"

func main() {
	// 定义切片
	var s1 []int
	var s2 []string

	fmt.Println(s1)
	fmt.Println(s2)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"aa", "bb", "cc"}
	fmt.Println(s1)
	fmt.Println(s2)

	// 由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s3 := a1[0:4] // 基于数组的一个切割，左闭右开
	fmt.Println(s3)
	// 切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
	fmt.Printf("len(s3): %d, cap(s3): %d\n", len(s3), cap(s3))
	a1[0] = 100
	// 切片是引用类型，都指向了底层的一个数组
	fmt.Println(s3)
}
