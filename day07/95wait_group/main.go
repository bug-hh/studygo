package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func task(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i)
	}
	// 等待 wg 的计数器减为零
	wg.Wait()
}
