package main

import "fmt"

func main() {

	// We start with a simple if statement. if statements evaluate boolean values
	// and will execute a branch of code if one is true or false
	if true == true {
		fmt.Println("True is true")
	} else { // Go also supports else and else if statements
		fmt.Println("True is not false")
	}

	// Go has one type of loop. The for loop. You can do typical c-style
	// iteration with a for loop
	for i := 0; i < 20; i++ {
		fmt.Println("Go!")
	}

	// The for loop can be used in multiple ways, for instance, to loop
	// indefinitely, simply use for {  }
	for {
		fmt.Println("This will go forever")
		// You can also break to stop the execution of the infinite loop
		break
	}

	// A more hybrid approach is to have the for loop evaluate a boolean value on
	// every iteration
	j := true
	for j {
		fmt.Println("This should happen once")
		// setting j to false should stop the loop
		j = false
	}

}
