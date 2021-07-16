package main

import "fmt"

func main() {
	dic := map[int]string{}
	dic[0] = "a"

	if v, ok := dic[0]; ok {
		fmt.Println(v)
	}
}
