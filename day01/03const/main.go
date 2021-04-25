package main

import "fmt"

const pi = 3.14

const (
	a = 10
	b = 20
	c = 30
)

// 这里 d,e,f 都等于 100
const (
	d = 100
	e
	f
)

const (
	aa = iota // 0
	bb        // 1
	cc        // 2
	dd        // 3
)

const (
	aaa = iota // 0
	bbb = 100
	ccc = iota // 2
	ddd = iota // 3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(ccc)
	fmt.Println(ddd)
}
