package main

import "fmt"

func main() {
	mySlice1 := []string{"I", "am", "stupid", "and", "weak"}
	fmt.Println("mySlice1", mySlice1)

	for index, c := range mySlice1 {
		if string(c) == "stupid" {
			mySlice1[index] = "smart"
		} else if string(c) == "weak" {
			mySlice1[index] = "strong"
		}
	}
	fmt.Println("mySlice1", mySlice1)

}
