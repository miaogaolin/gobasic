package unittest

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCache(t *testing.T) {
	content, err := ioutil.ReadFile("testdata/1.txt")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(content))
}
