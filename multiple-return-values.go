package main

import "fmt"

func vals() (int, int) { // Second parenthesis pair signifies multi-return
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
