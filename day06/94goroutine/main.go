package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i) // 开启一个单独的 goroutine 去单独执行 hello 函数
	}
	fmt.Println("main")     // main 函数结束了，那么由 main 函数开启的 goroutine 也都结束了
	time.Sleep(time.Second) // 让 main 函数晚点结束
}
