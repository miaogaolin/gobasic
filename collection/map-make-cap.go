package main

import "fmt"

func main() {
	dic := make(map[int]string, 10)
	dic[1] = "lao"
	fmt.Println("dic长度:", len(dic))
}
