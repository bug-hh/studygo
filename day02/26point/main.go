package main

import "fmt"

func main() {
	var p1 *int
	fmt.Println(p1)
	var p2 *int = new(int)
	fmt.Println(p2)
	fmt.Println(*p2)
	*p2 = 100
	fmt.Println(*p2)
}
