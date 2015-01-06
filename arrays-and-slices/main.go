package main

import "fmt"

func main() {
	// Go provides a very primitive array type. It is pretty close to the metal.
	// an array in Go is a numbered sequence of elements of a specific length.
	var nums [5]int
	fmt.Println("empty:", nums)

	nums[4] = 100
	fmt.Println("set:", nums)
	fmt.Println("get:", nums[4])

	// Arrays can be really really fast, but they arent super flexible. This is why
	// the slice abstraction was built.
	// Slices are a layer on top of arrays that make them easier to work with.
	// I like declaring slices like this
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("empty:", ints)

	// You can append items to a slice. This does not modify the slice value, but
	// returns a new slice instead.  keep in mind that this does not copy the
	// elements to the new slice, it manages all of the memory for you
	ints = append(ints, 6)
	fmt.Println("appended:", ints)

	// you can even use ranges to select a part of a slice. lets select different indexes
	fmt.Println("0-2:", ints[:2])
	fmt.Println("2-4:", ints[2:4])
	fmt.Println("4-6:", ints[4:])

	// slices can also be looped over in a special way using the range keyword
	for key, val := range ints {
		fmt.Println(key, val)
	}
}
