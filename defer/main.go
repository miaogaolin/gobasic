package main

import (
	"fmt"
	"io"
	"os"
)

func Fun1() {
	fmt.Println("猜我啥时候输出？")
}

// srcName 路径文件拷贝到 dstName 路径文件
func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	defer func() {
		fmt.Println("defer")
	}()
	defer Fun1()
	fmt.Println("main")
	panic("panic")
}
