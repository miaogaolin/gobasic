package main

import (
	"encoding/json"
	"fmt"
)

type Tag struct {
	Name string `json:"name"`
}

func main() {
	t := Tag{"tag"}
	b, _ := json.Marshal(t)
	fmt.Println(string(b))
}
