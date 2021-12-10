package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	select {
	case v, ok := <-ch1:
		if ok {
			fmt.Println("ch1通道", v)
		}
	case v, ok := <-ch2:
		if ok {
			fmt.Println("ch2通道", v)
		}
	}
}
