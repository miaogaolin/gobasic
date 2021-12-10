package unittest

import (
	"testing"
)

func TestAdd(t *testing.T) {
	excepted := 5
	actual := Add(2, 3)
	if excepted != actual {
		t.Errorf("excepted：%d, actual:%d", excepted, actual)
	}
}

func TestAdd1(t *testing.T) {
	excepted := 5
	actual := Add(2, 3)
	if excepted != actual {
		t.Errorf("case1：excepted：%d, actual:%d", excepted, actual)
	}

	excepted = 10
	actual = Add(0, 10)
	if excepted != actual {
		t.Errorf("case2：excepted：%d, actual:%d", excepted, actual)
	}
}

func TestAddTable(t *testing.T) {
	type param struct {
		name                 string
		num1, num2, excepted int
	}

	testCases := []param{
		{name: "case1", num1: 2, num2: 3, excepted: 5},
		{name: "case2", num1: 0, num2: 10, excepted: 10},
	}
	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			actual := Add(v.num1, v.num2)
			if v.excepted != actual {
				t.Errorf("excepted:%d, actual:%d", v.excepted, actual)
			}
		})
	}
}
