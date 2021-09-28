package main

import "fmt"

func SetCountry(countries map[string]string) {
	countries["china"] = "中国"
}

func main() {
	c := make(map[string]string)
	SetCountry(c)
	fmt.Println(c)
}
