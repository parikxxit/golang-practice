package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- 2
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	select {
	case val := <-ch1:
		fmt.Println(val)
	case val := <-ch2:
		fmt.Println(val)
	case <-ctx.Done():
		fmt.Println("Time out")
	}
}
