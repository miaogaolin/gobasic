package main

import "fmt"

func main() {
	// 方式一: 迭代计数
	nums := [...]int{3, 2, 1, 4}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// 方式二：for-range
	for i, v := range nums {
		fmt.Println("索引：", i, " 值", v)
	}
}
