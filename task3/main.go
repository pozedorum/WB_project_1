package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func worker(workerID int, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()

	for res := range jobs {
		fmt.Printf("worker %d say %d\n", workerID, res)
	}
}

func fillChan(jobs chan<- int) {
	for i := range 30 {
		jobs <- i
	}
	close(jobs)
}

func main() {
	jobs := make(chan int)

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <workers_count>")
		os.Exit(1)
	}
	workersCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: workers_count must be a number")
		os.Exit(1)
	}

	go fillChan(jobs)
	var wg sync.WaitGroup
	for workerID := range workersCount {
		wg.Add(1)
		go worker(workerID, &wg, jobs)
	}
	wg.Wait()
}
