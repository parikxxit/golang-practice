package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https:/api.github.com/user/parikxxit")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

}
