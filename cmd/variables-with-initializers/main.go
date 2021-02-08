package main

import "fmt"

var i, j int = 1, 2

// Variables with initializers (https://tour.golang.org/basics/9)
// A var declaration can include initializers, one per variable.
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
