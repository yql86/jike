package main

import "fmt"

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	fmt.Println("mySlice:", mySlice)
	fullSlice := myArray[:]
	fmt.Println("fullSlice:", fullSlice)
	remove3rdItem := deleteItem(fullSlice, 2)
	fmt.Println("remove3rdItem: ", remove3rdItem)
}

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
