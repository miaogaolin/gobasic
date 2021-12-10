package testcover

import "testing"

func TestGetSex(t *testing.T) {
	excepted := "男"
	actual := GetSex(1)
	if actual != excepted {
		t.Errorf("excepted：%s, actual:%s", excepted, actual)
	}
}
