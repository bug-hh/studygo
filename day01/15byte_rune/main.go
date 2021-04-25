package main

import "fmt"

// byte 和 rune 类型
// Go 语言中，为了处理非 ASCII 码类型的字符，定义了新的 rune 类型
func main() {
	// s := "hello, 彭豪辉，펑하오후이"
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%c\n", s[i])
	// }

	// for _, c := range s {
	// 	fmt.Printf("%c\n", c)
	// }

	s2 := "白萝卜"
	s3 := []rune(s2) // 将字符串强制转换成了一个 rune 切片
	s3[0] = '红'
	fmt.Printf("%s\n", string(s3)) // 将 rune 切片强制转换成 string
}
