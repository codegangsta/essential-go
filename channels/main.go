package main

import (
	"fmt"
	"time"
)

func main() {
	q := make(chan string)

	for i := 1; i <= 5; i++ {
		go worker(i, q)
	}

	for i := 0; i <= 500; i++ {
		q <- fmt.Sprintf("Message %v", i)
	}

	time.Sleep(1 * time.Second)
}

func worker(id int, queue chan string) {
	for msg := range queue {
		fmt.Printf("[worker %v] Doing work on '%v'\n", id, msg)
	}
}
