package main

import "fmt"

func main() {
	nums := [...]int{3, 2, 1, 4}
	copy := nums
	copy[1] = 3
	fmt.Println("nums:", nums)
	fmt.Println("copy:", copy)
}
