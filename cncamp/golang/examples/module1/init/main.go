package main

import (
	_ "./a"
	_ "./b"
	"fmt"
)

func init() {
	fmt.Println("main init")
}
func main() {

}
