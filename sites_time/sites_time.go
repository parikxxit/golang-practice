package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://apple.com",
		"https://facebook.com",
		"https://parikshit.dev",
	}
	var ws sync.WaitGroup
	ws.Add(len(urls))
	for _, url := range urls {
		url := url
		go func() {
			defer ws.Done()
			siteTime(url)
		}()
	}
	ws.Wait()
}

func siteTime(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Got Error: %#v", err)
		return
	}
	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Printf("Got error while copyaing: %#v", err)
	}
	duration := time.Since(start)
	log.Printf("Time take for url: %s to copy is %v", url, duration)
}
