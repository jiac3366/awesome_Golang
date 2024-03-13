package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type T struct {
		F1 int `json:"f_1"`
		F2 int `json:"f_2,omitempty"`
		F3 int `json:"f_3,omitempty"`
		F4 int `json:"-"`
	}
	t := T{1, 0, 2, 3}
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b) // {"f_1":1,"f_3":2}
}

// struct tag: https://go.dev/wiki/Well-known-struct-tags
// json struct tag: https://pkg.go.dev/encoding/json#Marshal
