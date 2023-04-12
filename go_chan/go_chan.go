package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine") // it will not print
	fmt.Println("main")
	/*
		for i := 0; i < 3; i++ {
			//BUG: it will always print the same value
			go func() {
				fmt.Println(i)
			}()

		}
			Output with above code
			main
			3
			3
			goroutine
			3
	*/
	//FIX1: Use a parameter
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	//FIX2 Use a Use shadow variable in loop
	for i := 0; i < 3; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}
	ch := make(chan string)
	/*
		ch <- "hi"  // send
		msg := <-ch // reveice
		fm, t.Println(msg)
	*/
	// Above commented code will cause deadlock as we are sending hi to the main go routine and no other is receving it so usually other go routine generate and main go routine consume
	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("got message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()
	for msg := range ch {
		fmt.Println("Got msg:", msg)
	}
	msg := <-ch // Channel is closed
	fmt.Printf("value of closed channel is #%v", msg)
	var ok bool
	msg, ok = <-ch
	if !ok {
		fmt.Printf("Channel is closed\n")
	} else {
		fmt.Println(msg)
	}
	time.Sleep(10 * time.Millisecond)
}
