package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNotFound = errors.New("not found")

var ErrHuman = fmt.Errorf("%s不符合我们人类要求", "老苗")

type ErrorPathNotExist struct {
	Filename string
}

func (*ErrorPathNotExist) Error() string {
	return "文件路径不存在"
}

var ErrNotExist error = &ErrorPathNotExist{
	Filename: "./main.go",
}

func main() {
	content, err := LoadConfig()
	if err != nil {
		if v, ok := err.(*FileEmptyError); ok {
			fmt.Println("Filename:", v.Filename)
		}
		log.Fatal(err)
	}

	fmt.Println("内容：", content)
}
