package main

import "fmt"

// 返回值没有起名字，必须显示的 return ret
func sum2(x int, y int) int {
	ret := x + y
	return ret
}

func sum1(x int, y int) (ret int) {
	ret = x + y
	// 因为我给返回值取了名字，所以可以只写个 return，函数依然会返回 ret 的值
	return
}

// 当多个参数具有相同类型，可以省略非最后一个参数类型
func sum3(x, y, z int, a, b, c string) int {
	return x + y + z
}

// 可变长参数
// 可变长参数必须放在函数参数最后
func f1(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y 可以传多个，也可以不传
}

func main() {
	s := sum2(2, 2)
	fmt.Println(s)

	f1("哈哈", 1, 2, 3, 4, 5, 6)
	f1("呵呵")

}
