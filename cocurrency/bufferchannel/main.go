package main

import "fmt"

func main() {
	bufferChan := make(chan string, 3)
	bufferChan<-"a"
	bufferChan<-"b"
	bufferChan<-"c"

	fmt.Println(<-bufferChan)
}
