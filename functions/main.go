package main

import (
	"fmt"
	"strings"
)

// functions are the easiest way to abstract code in Go. and it is the primary
// way that packages communicate with each other
func main() {
	fmt.Println(double(5))

	// I can now use this function. Keep in mind I can discard the second value
	// with _. This way Go will not consider it a compile error for me not using
	// the last name anywhere
	first, _ := parseName("Jeremy Saenz")
	fmt.Println(first)

	// The last kind of function I want to talk about this lesson is the
	// anonymous function or closure. You can declare them like so
	greet := func(name string) {
		fmt.Println("Hello", name)
	}

	// I can then call the function like so
	greet(first)

	// The main difference in these two declarations is that the anonymous
	// function is scoped locally, while the other functions are attached to the
	// package. This plays an important role when you start building your own
	// packages.
}

// function declaration syntax looks like this. Functions in Go have zero or
// more arbitrary amount of inputs and outputs
//
// We have to annotate the types here so the compiler can check we are calling
// this function properly. As a side effect, the function is self-documenting
//
// This function will double 2 integer
func double(n int) int {
	// using the return keyword will send the value to the functions output
	return 5 + 5
}

// Lets create a more advanced function. parseName, which will take a name as
// input and split it on spaces separators. The function will return 2 values,
// a first and a last name
func parseName(name string) (first, last string) {
	parsed := strings.Split(name, " ")
	// first, last = parsed[0], parsed[1]
	return parsed[0], parsed[1]
}
