package pkgB

import (
	"fmt"

	"github.com/miaogaolin/gobasic/pkgA"
)

var name string

func init() {
	name = "B"
}

func PrintName() {
	fmt.Println(name)
	pkgA.PrintName()
}
