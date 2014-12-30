package main

import "fmt"

// if you are not familiar with them, pointers can feel pretty daunting.
// Fortunately, Go prevents you from doing some terrible things with pointers
// that you can do in C/C++ and that is a great thing.
//
// The basics idea of a pointer is that it allows functions to reference and
// modify the variable you pass into it. The alternative is a Pass by Value,
// which means the value is copied into the function instead of being pointed
// to by a pointer.
//
// Let's play with some examples
func main() {
	num := 5
	double(num)
	// hm. That doesn't seem to work. This is because num is being copied into
	// the double function. Lets write another function that accepts a pointer to
	// an int instead of an int
	fmt.Println(num)

	// To convert a value to a pointer to that value, use &
	triple(&num)

	// Now num should be 15
	fmt.Println(num)
}

// Lets create another double function, but instead of returning a value, lets
// modify the value
func double(n int) {
	n = n * 2
}

// a pointer to a type is denoted with a *
func triple(n *int) {
	*n = *n * 3
}
