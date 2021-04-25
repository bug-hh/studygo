package main

import "fmt"

func main() {
	s := []string{"北京", "深圳", "上海"}
	// 调用 append 函数，必须用原来的切片变量接收返回值
	s = append(s, "长沙")
	fmt.Println(s)
	s = append(s, "武汉", "成都")
	fmt.Println(s)
	s2 := []string{"杭州", "广州", "苏州"}
	s = append(s, s2...) // ... 表示把 s2 每个元素单独拆开再添加
	fmt.Println(s)

	s3 := make([]string, 3, 3)
	s4 := s     // 赋值
	copy(s3, s) // 拷贝
	fmt.Println(s, s3, s4)
}
