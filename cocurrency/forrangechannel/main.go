package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 写通道
func write(data chan<- int)  {
	for i := 0; i < 10; i++ {
		data<-i
	}

	close(data)
	wg.Done()
}

// 读通道
func read(data <-chan int)  {
	for d := range data{
		fmt.Println(d)
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	dataChan := make(chan int)

	go write(dataChan)
	go read(dataChan)

	wg.Wait()
}
