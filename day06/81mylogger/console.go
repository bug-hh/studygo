package mylogger

import (
	"fmt"
	"time"
)

type Logger struct {
}

func newLogger() Logger {
	return Logger{}
}

func (l Logger) Debug(msg string) {
	now := time.Now()
	fmt.Printf("[%s] %s", now.Format("2006-01-02 15:04:05"), msg)

}
