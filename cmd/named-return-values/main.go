package main

import "fmt"

/*** notice (x, y int) as opposed to (string, string) in previous examples ***/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Named return values (https://tour.golang.org/basics/7)
// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
// These names should be used to document the meaning of the return values.
// A return statement without arguments returns the named return values. This is known as a "naked" return.
// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
func main() {
	fmt.Println(split(17))

	/*
	Multiple-value split(...) (int, int) in single-value context
	fmt.Printf("%s", split(85))
	fmt.Printf("%s, %s", split(85))
    */

	x, y := split(35)
	fmt.Printf("%d, %d\n", x, y)
}
