package main

import "fmt"

func main() {
	// div(0,1) panics
	fmt.Println(safeDiv(0, 1)) // recovers from panic
	fmt.Println(div(10, 3))
}
func div(a, b int) int {
	return a / b
}

func safeDiv(a, b int) (ans int, err error) {
	// ans and err are local variable
	if e := recover(); e != nil {
		// e is of type any not error
		err = fmt.Errorf("%v", e)
	}
	return a / b, nil
}
