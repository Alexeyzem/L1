//Реализовать все возможные способы остановки выполнения горутины.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("")
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Gorutine Stop")
				return
			default:
				fmt.Println("Gorutine work...")
			}
		}
	}()
	time.Sleep(time.Second)
	close(quit)

	ctx, closeCtx := context.WithCancel(context.Background())
	ch := make(chan int)
	go work(ctx, ch)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	closeCtx()
	time.Sleep(1 * time.Second)
}
func work(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop gorutine")
			return
		case num := <-ch:
			fmt.Println(num)
		}
	}
}
