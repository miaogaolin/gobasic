package main

import "fmt"

func main() {
	str := "我爱china"
	for i, c := range str {
		fmt.Printf("位置：%d, 字符：%c\n", i, c)
	}
}
