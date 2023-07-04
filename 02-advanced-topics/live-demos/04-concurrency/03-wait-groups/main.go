package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// Decrease the counter of the WaitGroup when the worker finishes
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// Simulate some work
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// Create a WaitGroup
	var wg sync.WaitGroup

	// Launch multiple workers
	for i := 1; i <= 3; i++ {
		// Increase the counter of the WaitGroup for each worker
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers have finished")
}
