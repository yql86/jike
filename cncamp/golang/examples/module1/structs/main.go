package main

import "reflect"

type MyType struct {
	Name string `json:"name"`
}

type TypeA struct {
	P1 string
}

type TypeB struct {
	P2 string
	TypeA
}

func main() {
	mt := MyType{Name: "test"}
	println(mt.Name)
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	println(name.Tag.Get("json"))

	tb := TypeB{P2: "p2", TypeA: TypeA{P1: "p1"}}
	println(tb.P1)
}
