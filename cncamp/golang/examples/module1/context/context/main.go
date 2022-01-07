package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseContext := context.Background()
	//ctx := context.WithValue(baseContext, "a", "b")
	//go func(c context.Context) {
	//	fmt.Println(c.Value("a"))
	//}(ctx)
	timeoutCtx, cancel := context.WithTimeout(baseContext, time.Second)
	defer cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)

	select {
	case <-timeoutCtx.Done():
		time.Sleep(time.Second)
		fmt.Println("main process exit")
	}
}
