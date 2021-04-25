package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	defer fileObj.Close()
	if err != nil {
		fmt.Println("打开文件出错")
		return
	}

	fileObj.Write([]byte("abc"))
	fileObj.WriteString("aaaa")
}
