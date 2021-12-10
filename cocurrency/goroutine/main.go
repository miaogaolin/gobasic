package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func PrintA()  {
	fmt.Println("A")
	wg.Done()
}

func main() {
	wg.Add(1)
	go PrintA()
	wg.Wait()
	fmt.Println("main")
}
