package main

import "fmt"

func main() {
	// numsCopy 长度小于 nums
	nums := []int{1, 2, 3}
	numsCopy := make([]int, 2)
	// 前面两个元素 1 和 2 被复制
	copy(numsCopy, nums)
	fmt.Println("numsCopy(小于):", numsCopy)

	// numsCopy 长度大于 nums
	numsCopy = []int{4, 5, 6, 7}
	// 4,5,6 会被覆盖，保留 7
	copy(numsCopy, nums)
	fmt.Println("numsCopy(大于):", numsCopy)
}
