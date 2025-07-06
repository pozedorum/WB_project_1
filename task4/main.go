package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

/*
Я попробовал останавливать рабочих через дополнительный канал,
при получении SIGINT я закрывал канал, это считывалось рабочим и горутина завершалась,
но так как сообщение было одно и оно обрабатывалось, остальные рабочие его не получали.
Пришлось использовать контекст, так как он принудительно завершает горутины, когда человек нажимает ctrl+C
*/

func worker(ctx context.Context, workerID int, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()

	for {
		select {
		case res, ok := <-jobs:
			if !ok {
				return
			} else {
				fmt.Printf("worker %d say %d\n", workerID, res)
			}
		case <-ctx.Done():
			fmt.Printf("worker %d is ended\n", workerID)
			return
		}
	}
}

func fillChan(jobs chan<- int) {
	for i := range 30 {
		jobs <- i
		time.Sleep(time.Second)
	}
	close(jobs)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <workers_count>")
		os.Exit(1)
	}

	jobs := make(chan int)
	done := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	workersCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: workers_count must be a number")
		os.Exit(1)
	}

	go fillChan(jobs)
	var wg sync.WaitGroup
	for workerID := range workersCount {
		wg.Add(1)
		go worker(ctx, workerID, &wg, jobs)
	}
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-cancelChan:
		fmt.Println("\nReceived interrupt signal")
	case <-done:
		fmt.Println("\nWork completed successfuly")
	}

	cancel()
	wg.Wait()
}
