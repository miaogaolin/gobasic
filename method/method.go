package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func (p People) GetName() string {
	return p.Name
}

func (p People) SetName(name string) string {
	p.Name = name
	return name
}

// 内置类型
type Num int

func (n Num) String() string {
	return fmt.Sprintf("%d", n)
}
func (n *Num) Set(num Num) {
	n = &num
}

type M map[string]string

func (m M) SetKey(key, val string) {
	(m)[key] = val
}

func main() {
	var n Num = 2
	n.Set(3)
	fmt.Println(n)

	p := &People{Name: "苗"}
	p.SetName("潇洒哥")
	fmt.Println(p.Name)

	m := make(M)

	m.SetKey("name", "苗")

	fmt.Println(m)
}
