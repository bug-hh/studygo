package main

import "fmt"

func main() {
	// age 变量只在 if 条件语句作用域中生效
	if age := 19; age > 18 {
		fmt.Println("澳门首家赌场开业了")
	} else {
		fmt.Println("回去写作业")
	}

	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()

	// 第二种
	j := 0
	for ; j < 10; j++ {
		fmt.Printf("%4d", j)
	}
	fmt.Println()

	// 第三种
	z := 0
	for z < 10 {
		fmt.Printf("%4d", z)
		z++
	}
	fmt.Println()

	// 无限循环
	// for {
	// 	fmt.Println("123")
	// }

	// for range
	s := "hello彭豪辉"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	// go 语言中的 switch 语句可以省略 break
	t := 2
	switch t {
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	case 4:
		fmt.Println("D")
	default:
		fmt.Println("S")
	}

	breakLable()
}

func breakLable() {
BREAKLABEL:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKLABEL
			}
			fmt.Printf("%d %d\n", i, j)
		}
	}
	fmt.Println("...")
}
