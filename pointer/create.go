package main

import "fmt"

func main() {
	var p *int
	var num int = 11

	p = &num
	fmt.Println(p)
	fmt.Println(*p)

	var empty *int
	fmt.Println(empty == nil)
}
