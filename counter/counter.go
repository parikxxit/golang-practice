package main

import "sync"

func main() {
	count := 0
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	// The above code will have race condition and give different result everytime u run it
}
