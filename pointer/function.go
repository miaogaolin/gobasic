package main

import "fmt"

func UpdateNum(num *int) {
	*num = 2
}

func main() {
	n := 1
	UpdateNum(&n)
	fmt.Println(n)
}
