package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var b chan int
	fmt.Println(b)     // nil
	b = make(chan int) // 不带缓冲区通道初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Printf("从通道道成功取出 %d\n", x)
	}()
	b <- 10
	fmt.Println("往通道中放入 10")
	b = make(chan int, 16) // 带缓冲区通道初始化
	fmt.Println(b)
	wg.Wait()
}
