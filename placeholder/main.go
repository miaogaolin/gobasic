package main

import (
	"errors"
	"fmt"
)

type Example struct {
	Content string
}

var data = Example{Content: "例子"}

type str string

func main() {
	fmt.Printf("%v", data)

	fmt.Printf("%v", errors.New("我错了"))

	fmt.Printf("%+v", data)

	fmt.Printf("%#v", data)
}
