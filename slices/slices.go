package main

import "fmt"

func main() {
	fmt.Println(concatinate([]string{"abs", "Def"}, []string{"some thing", "new"}))
}

func concatinate(s1, s2 []string) []string {
	return append(s1, s2...)
}
