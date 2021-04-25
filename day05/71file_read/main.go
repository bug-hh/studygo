package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileObj, err := os.Open("./main.go")
	defer fileObj.Close()

	if err != nil {
		fmt.Println("读文件出错")
		return
	}

	var ret [128]byte
	for {
		n, readError := fileObj.Read(ret[:])
		if readError == io.EOF {
			fmt.Println("文件读完了")
			return
		}

		if readError != nil {
			fmt.Println("读文件出错")
			return
		}
		fmt.Printf("读了 %d 个字节\n", n)
		fmt.Println(string(ret[:n]))
		if n < 128 {
			return
		}
	}
}
