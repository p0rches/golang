package main

import "fmt"

func zero(val int) {
	val = 0
}

func zeropt(iptr *int) {
	*iptr = 0
}

func main() {

	var i int = 1

	fmt.Println(i)
	fmt.Println(i)
	fmt.Println()
	fmt.Println()
}
