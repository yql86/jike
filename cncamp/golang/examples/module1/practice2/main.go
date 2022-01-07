package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	messages := make(chan int, 10)
	defer close(messages)

	//consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			i := <-messages
			fmt.Println("get: ", i)
		}
	}()

	//producer
	putTicker := time.NewTicker(1 * time.Second)
	for _ = range putTicker.C {
		i := rand.Intn(100)
		fmt.Println("put: ", i)
		messages <- i
	}

}
