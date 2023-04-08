package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/parikxxit")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error := %s", resp.StatusCode)
	}
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("error: cant copy %s", err)
	}
}
