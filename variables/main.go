// Statically compiled languages usually get a bad rap because of how verbose
// they usually are in declaring which variables are what types. Go keeps
// itself succinct by having a couple different ways to declare variables. One
// of which makes it feel very much like a dynamic language rather than a
// static one. Let's look at the multiple ways variables are declared.
package main

import "fmt"

// So far we talked about local variables. What about package-wide variables?
// We can declare package wide varables using the "var" syntax outside of a
// function.
//
// If you have multiple package level variables together, it is idiomatic to
// group them like so.
var (
	// Case is significant in Go. If a identifier (Function, Type, Variable, or
	// Const) is uppercase then it is considered exported, or publicly available
	// for other packages to use when they import said package.
	//
	// This is why all the functions we have been calling so far from the fmt
	// package have been uppercase, it is because they have to
	Version = "0.0.1"
	// Consequently, anything not capitalized is only visible to the package it
	// is contained in. This pattern at first feels a little weird and
	// restrictive, but it actually is an excellent communication mechanism. You
	// can easliy see what is private and what is public.
	//
	// Since we are in the main package right now, and nobody can import the main
	// package, what we declare public and private does not really matter in the
	// long run.
	author = "@codegangsta"
)

// Constants can be declared the same way that variables are. And it's also
// idiomatic to group them together
const (
	CCVisa       = "Visa"
	CCMasterCard = "MasterCard"
	CCAmex       = "American Express"
)

func main() {
	// The first (and most verbose) way to declare and initialize a variable is
	// this way. Using the var keyword paired with the identifier and the type.
	// Finally after the equals sign we can initialize the variable with a value.
	//
	// Of course, these are variables, we can also assign then different values.
	// If the variable isn't explicitly assigned a value it will still be
	// initialized with a zero value. Say goodbye to uninitialized values from C
	var greeting string = "Hello world!"
	// We can also declare multiple variables in one line and assign them values.
	var a, b, c int = 1, 2, 3
	fmt.Println(greeting, a, b, c)

	// The second way to declare variables is to simply drop the type annotation.
	// Keep in mind that this only works if you are assigning a value to the
	// variable on the same line. Go needs some way to infer what the variable
	// is.
	var name = "Jeremy Saenz"
	// This works nicely with multiple variables as well. d will be an int, e an
	// float, and f a string
	var d, e, f = 1, 2.0, "Three"
	fmt.Println(name, d, e, f)

	// The final (and most idiomatic) way of declaring variables is using the
	// Short Variable Declaration syntax. Drop the var and add a ":" before the
	// "="
	//
	// Now our declaration for variables is just as succinct (and way more safe)
	// than a dynamic language.
	course := "Essential Go"
	// We can do the same thing with multiple assignment as well.
	x, y, z := 1, 2, 3
	fmt.Println(course, x, y, z)
}
