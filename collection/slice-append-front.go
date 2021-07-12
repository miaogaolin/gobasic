package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	nums = append([]int{4}, nums...)

	fmt.Println(nums)
}
