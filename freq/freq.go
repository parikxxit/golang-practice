package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

// Q what is the most common word in sherlock.txt aka word freq
func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer file.Close()

	word, cnt := mostCommon(file)
	fmt.Println(word, cnt)
}

// "Who's on first?" -> [who s on first]
var WordRe = regexp.MustCompile(`[a-zA-Z]+`)

// TODO change it to return n most common freq
func mostCommon(r io.Reader) (string, int) {
	freq, err := wordFrequency(r)
	if err != nil {
		log.Fatalf("error wile calling wordFreq: %v", err)
		return "", 0
	}
	return maxWordFreq(freq)
}

func maxWordFreq(freqMap map[string]int) (string, int) {
	max, word := 0, ""
	for k, v := range freqMap {
		if v > max {
			max = v
			word = k
		}
	}
	return word, max
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freq := make(map[string]int) // word -> count
	for s.Scan() {
		words := WordRe.FindAllString(s.Text(), -1)
		for _, w := range words {
			freq[w]++
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return freq, nil
}
