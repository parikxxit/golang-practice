package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/parikxxit")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error := %s", resp.StatusCode)
	}
	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// log.Fatalf("error: cant copy %s", err)
	// }
	var r Replay
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("Unabl decode %s", err)
	}
	fmt.Printf("%#v\n", r)
}

type Replay struct {
	Name         string
	Public_Repos int
}
