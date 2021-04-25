package main

import "fmt"

func main() {
	var arr [3]int
	arr = [3]int{1, 2, 3}
	fmt.Println(arr)
	// 初始化方式2
	// 使用 ... 自动推断数组元素个数
	arr2 := [...]int{1, 2, 3, 4}
	fmt.Println(arr2)
	// 给指定的索引对应的数初始化, 第 0 个数初始化成 2，最后一个数初始化成 100，其余默认为 0
	arr3 := [5]int{0: 2, 4: 100}
	fmt.Println(arr3)

	// 多维数组
	arr11 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(arr11)
}
