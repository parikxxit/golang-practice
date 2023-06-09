package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := sha1sum("http.log.gz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sig)
	sig, err = sha1sum("sha1.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sig)
}

// mimic operation cat http.log.gz | gunzip | sha1sum
func sha1sum(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var r io.Reader = file
	if strings.HasSuffix(filename, "gz") {
		r, err = gzip.NewReader(file)
		if err != nil {
			return "", err
		}
	}
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}
	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
