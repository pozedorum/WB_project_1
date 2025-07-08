package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func worker(workerID int, wg *sync.WaitGroup, counter *int32) {
	defer wg.Done()

	for range 10 {
		atomic.AddInt32(counter, int32(workerID))
	}
}

func main() {

	workersCount := 10
	var counter int32 = 0
	mustResult := 0
	var wg sync.WaitGroup
	for workerID := 1; workerID <= workersCount; workerID++ {
		wg.Add(1)
		go worker(workerID, &wg, &counter)
		mustResult += 10 * workerID
	}
	wg.Wait()
	fmt.Println("expected result:", mustResult)
	fmt.Println("goroutines result:", counter)

}

// До этого работал только с мьютексами, атомик тоже довольно интересен
