package main

import "fmt"

func main() {
	// 往回跳，当 i 不小于 2 时，结束跳转
	i := 0
UP:
	fmt.Println(i)
	if i < 2 {
		i++
		goto UP
	}

	// 往后跳，跳过 goto 与 DOWN: 之间的逻辑
	fmt.Println("你")
	goto DOWN
	fmt.Println("好")
DOWN:
	fmt.Println("呀")

}
