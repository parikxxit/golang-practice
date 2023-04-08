package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	user, repos, err := githubInfo("parikxxit")
	if err != nil {
		log.Fatalf("got error: %s", err)
	}
	fmt.Printf("Username: %s, repos: %d\n", user, repos)
}

func githubInfo(login string) (string, int, error) {
	URL := fmt.Sprintf("https://api.github.com/users/%s", login)
	resp, err := http.Get(URL)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, errors.New("not ok status code")
	}
	var r Replay
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("Unabl decode %s", err)
	}
	return r.Name, r.Public_Repos, nil
}

type Replay struct {
	Name         string
	Public_Repos int
}
