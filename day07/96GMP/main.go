package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("A: %d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("B: %d\n", i)
	}
}

func main() {
	// runtime.GOMAXPROCS(1)  默认CPU 逻辑核心数，默认跑满
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
