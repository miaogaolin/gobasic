package testatomic

import "testing"

func TestAddNumber(t *testing.T) {
	excepted := 2
	actual := AddNumber(excepted)
	if excepted != actual {
		t.Errorf("excepted：%d, actual:%d", excepted, actual)
	}
}
