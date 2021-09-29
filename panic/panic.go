package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("我捕捉的：", err)
			fmt.Println("我好好的")
		}
	}()
	panic("我是异常")
}
