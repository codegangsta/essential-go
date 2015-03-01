package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 2)

	go func() {
		time.Sleep(250 * time.Millisecond)
		done <- true
	}()

	go func() {
		time.Sleep(1000 * time.Millisecond)
		done <- true
	}()

	waiton("job 1", done)
	waiton("job 2", done)
}

func waiton(job string, done chan bool) {
	select {
	case <-done:
		fmt.Println(job + " Finished")
	case <-time.After(500 * time.Millisecond):
		fmt.Println(job + " Timeout")
	}
}
