package main

import (
	"log"
	"time"
)

func commonLog() {
	for {
		log.Println("这是一条测试日志")
		time.Sleep(time.Second * 3) // 每隔 3 s 打一次日志
	}
}

func main() {
	commonLog()
}
