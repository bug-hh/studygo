package main

import "fmt"

func main() {
	var m map[string]int
	m = make(map[string]int, 10)
	m["aaa"] = 1
	m["bbb"] = 2

	// map 中，如果没有对应的 key，那么会打印 value 类型的默认零值
	fmt.Println(m["娜扎"]) // 0
	v, ok := m["娜扎"]     // 注意这里写 :=  而不是 =
	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Println(v)
	}

	// map 的遍历
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 只遍历 key
	for k := range m {
		fmt.Println(k)
	}

	// 只遍历 value
	for _, v := range m {
		fmt.Println(v)
	}

	// 删除
	delete(m, "aaa")
	fmt.Println(m)

	// map + 切片
	// 元素类型为 map 的切片
	var marr = make([]map[int]string, 10, 10) // 这里第二个参数长度一定要大于 0
	marr[0] = make(map[int]string, 10)        // 给 map 分配内存
	marr[0][0] = "abc"

	fmt.Println(marr)

	// 值为切片类型的 map
	var mm = make(map[int][]int, 10)
	mm[0] = []int{12, 13, 14}
	fmt.Println(mm)

}
