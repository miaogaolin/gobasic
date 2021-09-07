package main

import (
	"fmt"
)

type People struct {
	Name string
	Age  int
}

type Student struct {
	People
	Collect string
}

type RepeatStudent struct {
	People
	Collect string
	Name    string
}

type AnoStudent struct {
	People struct {
		Name string
		Age  int
	}
	Collect string
}

func main() {
	s1 := Student{
		People: People{
			Name: "老苗",
			Age:  18,
		},
		Collect: "不告诉",
	}

	s2 := Student{
		People{
			Name: "老苗",
			Age:  18,
		},
		"不告诉",
	}
	fmt.Println(s1, s2)

	r := RepeatStudent{
		People: People{Name: "老苗"},
		Name:   "潇洒哥",
	}
	fmt.Println(r.Name, r.People.Name)

	// 匿名
	ano := struct {
		Name string
	}{
		Name: "匿名",
	}
	fmt.Println(ano.Name)
}
