package main

import "fmt"

func main() {
	dic := make(map[int]string)
	dic[1] = "lao"
	dic[3] = "miao"
	fmt.Println("dic长度:", len(dic))
}
