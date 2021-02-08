package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

// Multiple results (https://tour.golang.org/basics/6)
// A function can return any number of results.
// The swap function returns two strings.
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
