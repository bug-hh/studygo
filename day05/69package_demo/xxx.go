package main

import (
	"fmt"

	// 一般来说包名和「包所在文件夹名」一致，如果不一致，可以给它取个别名
	calc "github.com/studygo/day05/69package"
)

func main() {
	r := calc.Add(10, 20)
	fmt.Println(r)
}
