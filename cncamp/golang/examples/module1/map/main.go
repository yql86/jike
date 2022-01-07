package main

import "fmt"

func main() {
	myMap := make(map[string]string, 10)
	myMap["key1"] = "value1"
	println("myMap: ", myMap["key1"])

	myFuncMap := map[string]func() int{
		"funcA": func() int {
			return 1
		},
	}
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	fmt.Println(f())

	value, exists := myMap["a"]
	if exists {
		println(value)
	} else {
		println("not exist")
	}

	for k, v := range myMap {
		println(k, v)
	}

}
