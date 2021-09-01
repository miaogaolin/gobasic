package pkgA

import "fmt"

var name string

func init() {
	name = "A"
}

func PrintName() {
	fmt.Println(name)
}
