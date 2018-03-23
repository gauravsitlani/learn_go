package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := message{"Alice", "Hello", 123123}
	b, _ := json.Marshal(m)
	fmt.Printf("%s\n", b)
	var ms message
	err := json.Unmarshal(b, &ms)
	bb, _ := json.Marshal(ms)
	fmt.Printf("%s", bb)
	fmt.Println(err)
}
