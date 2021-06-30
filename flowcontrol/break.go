package main

import "fmt"

func main() {
	// 默认跳出一层：当 i 为 5 时，跳出循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// 指定跳出：当 j 为 5 时，跳出两层循环
TWO:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break TWO
			}
			fmt.Println(i)
		}
	}
}
