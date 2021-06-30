package main

import "fmt"

func main() {
	switch num1, num2 := 1, 2; {
	case num1 >= num2:
		fmt.Println("num1 > num2")
	case num1 < num2:
		fmt.Println("num2 < num1")
	}

}
