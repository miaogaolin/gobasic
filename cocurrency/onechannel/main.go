package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 写通道
func write(data chan<- int)  {
	data<-520
	wg.Done()
}

// 读通道
func read(data <-chan int)  {
	fmt.Println(<-data)
	wg.Done()
}

func main() {
	wg.Add(2)

	dataChan := make(chan int)

	go write(dataChan)
	go read(dataChan)

	wg.Wait()
}
