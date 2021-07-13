package main

import "fmt"

func main() {
	// 将 nums 拷贝到 numsCopy
	nums := []int{1, 2, 3}
	numsCopy := make([]int, 2)
	copy(numsCopy, nums)
	numsCopy[0] = 2
	fmt.Println("nums:", nums)
	// 修改了 numsCopy，不会对 nums 产生影响
	fmt.Println("numsCopy:", numsCopy)
}
