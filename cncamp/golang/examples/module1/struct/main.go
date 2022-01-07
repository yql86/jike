package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City string `json:"city"`
}

type MyType struct {
	Name string `json:"name"`
	Address
}

func main() {
	mt := MyType{Name: "test", Address: Address{City: "hangzhou"}}
	b, _ := json.Marshal(mt)
	fmt.Println(string(b))
}
