package main

import (
	"fmt"
)

type People interface {
	Eat(thing string) error
	Drink(thing string) error
}

type LaoMiao struct {
	Name string
	Age  int
}

func (l LaoMiao) Eat(thing string) error {
	fmt.Println("在公司偷吃" + thing)
	return nil
}

func (l LaoMiao) Drink(thing string) error {
	fmt.Println("在公司偷喝" + thing)
	return nil
}

type LaoSun struct {
	Name string
	Age  int
}

func (l LaoSun) Eat(thing string) error {
	fmt.Println("在车上吃" + thing)
	return nil
}

func (l LaoSun) Drink(thing string) error {
	fmt.Println("在车上喝" + thing)
	return nil
}

type GouDan struct {
	Name string
	Age  int
}

func (l *GouDan) Eat(thing string) error {
	fmt.Println("你管我吃" + thing)
	return nil
}

func (l *GouDan) Drink(thing string) error {
	fmt.Println("你管我喝" + thing)
	return nil
}

func Run(p People) {
	thing1, thing2 := "桃子", "可乐"
	p.Eat(thing1)
	p.Drink(thing2)
}

type Student interface {
	Eat(thing string) error
	Drink(thing string) error
	Study()
}

type Empty interface{}

func main() {
	var p People = LaoMiao{}
	p.Eat("桃子")

	m := LaoMiao{}
	m.Eat("桃子")

	Run(LaoMiao{})
	Run(LaoSun{})

	Run(&LaoSun{})

	var sun *LaoSun = &LaoSun{}
	Run(sun)

	var people People

	// 将 People 类型转化为 LaoMiao 值类型
	people = LaoMiao{}
	val := people.(LaoMiao)
	fmt.Println(val)

	// 将 People 类型转化为 LaoMiao 指针类型
	people = &LaoMiao{}
	peo := people.(*LaoMiao)
	fmt.Println(peo)

	people = LaoMiao{}

	_, ok := people.(*LaoMiao)
	fmt.Println(ok)

	_, ok = people.(LaoMiao)
	fmt.Println(ok)

}
