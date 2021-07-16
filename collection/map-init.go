package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 2,
		"b": 3,
		"c": 4,
	}

	fmt.Println("b:", m["b"])
}
