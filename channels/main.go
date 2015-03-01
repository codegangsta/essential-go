package main

import (
	"fmt"
	"time"
)

func main() {
	jobcount := 100
	workercount := 5
	jobs := make(chan int)
	results := make(chan int)

	// Fire up the workers
	for i := 1; i <= workercount; i++ {
		go worker(i, jobs, results)
	}

	// Send a bunch of jobs over
	go func() {
		for i := 0; i < jobcount; i++ {
			jobs <- i
		}
	}()

	// Get the results from all of our jobs
	for i := 0; i < jobcount; i++ {
		<-results
	}
}

func worker(id int, jobs chan int, results chan int) {
	for jobid := range jobs {
		fmt.Printf("[worker %v] job %v\n", id, jobid)

		// Do some long term processing here
		time.Sleep(1000 * time.Millisecond)

		// Send the job id to the results
		results <- jobid
	}
}
