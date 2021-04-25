package main

import (
	"bufio"
	"fmt"
	"os"
)

func useScan() {
	var s string
	// 这种读取方式，输入中不能有空格
	fmt.Scanln(&s)
	fmt.Println(s)
}
func main() {
	var s string
	// 使用 bufio 读取带空格的输入
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Println(s)
	useScan()
}
