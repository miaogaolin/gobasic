package main

import (
	"encoding/binary"
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

	fmt.Printf("%T", data)

	fmt.Printf("%%")

	fmt.Printf("%t", true)

	fmt.Printf("%b", 4)

	fmt.Printf("%c", 0x82d7)

	fmt.Printf("%d,%d,%d", 10, 010, 0x10)
	fmt.Printf("|%5d|%-5d|%05d|", 1, 1, 1)

	fmt.Printf("%o,%o,%o", 10, 010, 0x10)
	fmt.Printf("%#o", 10)

	fmt.Printf("%q", 0x82d7)

	fmt.Println()
	fmt.Printf("%x, %#x", 13, 13)

	fmt.Println()
	fmt.Printf("%X, %#X", 13, 13)

	fmt.Println()
	fmt.Printf("%U", 0x82d7)

	fmt.Println()
	fmt.Printf("%#U", 0x82d7)

	fmt.Println()
	fmt.Printf("%b", 0.1)

	fmt.Println()
	fmt.Printf("%e", 10.2)

	fmt.Println()
	fmt.Printf("%E", 10.2)

	fmt.Println()
	fmt.Printf("%f", 10.2)

	fmt.Println()
	fmt.Printf("%.2f|%.2f", 10.232, 10.235)

	fmt.Println()
	fmt.Printf("%g|%g", 10.20, 1.200000+3.400000i)
	fmt.Println()
	fmt.Printf("%g|%g", 2e2, 2E2)

	fmt.Println()

	fmt.Printf("%.3g", 12.34)
	fmt.Println()

	fmt.Printf("%s|%s", "老苗", []byte("嘿嘿嘿"))
	fmt.Println()
	fmt.Printf("%q", "老苗")

	fmt.Println()
	fmt.Printf("%x|%X", "苗", "苗")

	fmt.Println()
	num := 2
	s := []int{1, 2}
	fmt.Printf("%p|%p", &num, s)

	fmt.Println()
	fmt.Printf("%+d|%+d", 2, -2)

	fmt.Println()
	fmt.Printf("%+q|%+q", "miao","苗")

	fmt.Println()
	fmt.Printf("%#q", "苗")
	fmt.Println()

	fmt.Printf("|% d|", 2)
	fmt.Println()

	fmt.Printf("%05s", "a")

	fmt.Println()
	fmt.Printf("%+05d", 1)

	fmt.Println()
	fmt.Printf("%.2s", "abc")

binary.Write()

}
