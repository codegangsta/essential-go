package main

import "fmt"

// a map is a typed key-value store in go. You can add items by specifying the
// key you want to associate it with. You can also look these items up later.
func main() {

	// like slices, maps can be initialized with the make builtin
	// This assigns to "ages" a map with strings as keys and integers as values
	age := make(map[string]int)

	// We can now start putting things in this map
	age["jeremy"] = 24
	age["jordie"] = 21
	age["josh"] = 27
	fmt.Println(age)

	// You can look up keys by using the square bracket syntax
	fmt.Println("Jeremy's age:", age["jeremy"])

	// You can remove keys using the delete builtin
	delete(age, "jeremy")
	fmt.Println(age)

	// Using the len keyword will give you the number of key/value pairs in the
	// map It should be 2 now
	fmt.Println(len(age))

	// You can also declare a map without use of make. I personally like this way
	// better
	m := map[string]int{
		"jeremy": 24,
		"jordie": 21,
		"josh":   27,
	}
	fmt.Println(m)

	// you can range over a map just like a slice
	for n, a := range m {
		fmt.Printf("%v is %v years old.\n", n, a)
	}
}
