package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	// 必须将成员变量首字母大写，这样才能保证在其他包中能够访问该成员
	// 使用 `` 指定在不同格式（json 等)的文件中，该成员的名字
	Name string `json: "name" db: "name" ini: "name"`
	Age  int    `json: "age"`
}

func main() {
	p1 := person{
		Name: "abc",
		Age:  18,
	}

	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err: %v\n", err)
	}

	fmt.Printf("%v\n", string(b))

	// 反序列化
	var p2 person
	s := `{"name":"abc","age":18}`
	json.Unmarshal([]byte(s), &p2)
	fmt.Printf("%#v\n", p2)

}
