package main

import (
	"fmt"
	"time"
)

func main() {
	// say("go")
	// go say("for it")

	go say("for it")
	say("go")
}

func say(message string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(message)
	}
}
