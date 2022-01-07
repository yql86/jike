package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("main")
	go lock()
	time.Sleep(5 * time.Second)
}

func lock() {
	//deadlock
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("lock: ", i)
	}
}
