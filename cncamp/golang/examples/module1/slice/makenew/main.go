package main

import "fmt"

func main() {
	mySlice1 := new([]int)
	mySlice2 := make([]int, 0)
	mySlice3 := make([]int, 10)
	mySlice4 := make([]int, 10, 20)
	fmt.Println("mySlice1:", &mySlice1)
	fmt.Println("mySlice2:", mySlice2)
	fmt.Println("mySlice3:", mySlice3)
	fmt.Println("mySlice4:", mySlice4)
}
