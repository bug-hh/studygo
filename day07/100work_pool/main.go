package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	jb   *job
	ssum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func generate(ch chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		ch <- newJob
		time.Sleep(time.Second)
	}
}

func calc(ch1 <-chan *job, ch2 chan<- *result) {
	defer wg.Done()
	for {
		task := <-ch1
		n := task.value
		var sum int64 = 0
		for n > 0 {
			sum += (n % 10)
			n /= 10
		}
		newResult := &result{
			jb:   task,
			ssum: sum,
		}
		ch2 <- newResult
	}

}

func main() {
	wg.Add(1)
	go generate(jobChan)
	wg.Add(24)
	// 开启 24 个 goroutine
	for i := 0; i < 24; i++ {
		go calc(jobChan, resultChan)
	}
	// 在主 goroutine 中打印结果
	for ret := range resultChan {
		fmt.Printf("value: %d, sum: %d\n", ret.jb.value, ret.ssum)
	}
	wg.Wait()

}
