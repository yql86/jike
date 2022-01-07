package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//fmt.Println("Hello world"

	//name := flag.String("name", "world", "specify the name you want to say hi")
	//flag.Parse()
	//fmt.Println("os args is:", os.Args)
	//fmt.Println("input parameter is:", *name)
	//fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	//fmt.Println(fullString)

	//var fullString string = "hello"
	//fmt.Println(fullString)
	//for i, c := range fullString {
	//	fmt.Println("out: ", i, string(c))
	//}

	//ret1, ret2 := passValue(1, 2)
	//fmt.Println("ret1=", ret1)
	//fmt.Println("ret2=", ret2)

	//sum(1, 2)
	//sum(1, 2, 4)
	//nums := []int{1, 2, 3, 4}
	//sum(0, nums...)

	//i := func(x, y int) int { return x + y }
	//println(i(2, 3))

	//ch := make(chan int)
	//go func() {
	//	fmt.Println("hello from goroutine")
	//	ch <- 0
	//}()
	//
	//i := <-ch
	//fmt.Println("receive: ", i)

	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10)
			fmt.Println("putting:", n)
			ch <- n
		}
		close(ch)
	}()

	fmt.Println("hello from main")
	for v := range ch {
		fmt.Println("receiving:", v)
	}
}

//func passValue(a, b int) (x, y int) {
//	x = a + 1
//	y = b + 2
//
//	return x, y
//}

//func sum(initV int, nums ...int) {
//	fmt.Println(nums, " ")
//	total := initV
//	for _, num := range nums {
//		total += num
//	}
//	fmt.Println(total)
//}
