package main

import "fmt"

func main() {
	// c := make(chan string)
	c := make(chan string, 2)
	c <- "Hello"
	c <- "world"

	fmt.Println(<-c)
	fmt.Println(<-c)
}
