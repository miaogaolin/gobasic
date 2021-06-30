package main

import "fmt"

func main() {
	// 不指定位置
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				fmt.Printf("%d+%d我没算\n", i, j)
				continue
				fmt.Println("没我的份")
			}
			fmt.Printf("%d+%d=%d\n", i, j, i+j)
		}
	}

	// 指定位置
LABEL:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 {
				fmt.Printf("%d+%d我没算\n", i, j)
				continue LABEL
				fmt.Println("没我的份")
			}
			fmt.Printf("%d+%d=%d\n", i, j, i+j)
		}
	}
}
