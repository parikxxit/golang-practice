package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G😀", 6)

	s := "G😀"
	fmt.Println("len: ", len(s))
	// code point = rune ~= unicode char
	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c is of type %T\n", r, r)
			// rune i.e int32
		}
	}

	b := s[0]
	fmt.Printf("%c is of type %T", b, b)
	// byte utf-8
}

func banner(text string, width int) {
	// padding := (width - len(text)) / 2 BUGGY code as it travers through byte
	padding := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
