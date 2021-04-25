package main

import (
	"fmt"
	"sync"
)

// 先往 channel 1 中存 100 个数
// 再从 channel 1 中取出 i，将 i * i 存入 channel 2

var wg sync.WaitGroup
var once sync.Once

func f1(ch chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		// 如果通道关闭了，那么 ok 为 false
		if !ok {
			break
		}
		ch2 <- x * x
	}
	// 确保某个操作只执行一次
	once.Do(func() {
		close(ch2)
	})
}

func main() {
	wg.Add(3)
	var ch1 = make(chan int, 100)
	var ch2 = make(chan int, 100)
	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)
	wg.Wait()
	for n := range ch2 {
		fmt.Println(n)
	}
}
