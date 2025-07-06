package main

import (
	"fmt"
	"time"
)

func fillChan(jobs chan<- int) {
	for i := 20; i > 0; i-- {
		jobs <- i
		time.Sleep(1 * time.Second)
	}
	close(jobs)
}

func getFromChan(n int, jobs <-chan int, done chan<- bool) {
	nTime := time.After(time.Duration(n) * time.Second)
	for {
		select {
		case num, ok := <-jobs:
			if !ok {
				fmt.Println("\nWork completed successfuly")
				done <- true
				return
			}
			fmt.Println(num)
		case <-nTime:
			fmt.Println("\nWork canceled by timeout")
			done <- true
			return
		}
	}
}

func main() {
	var n int
	jobs := make(chan int)
	done := make(chan bool)
	fmt.Print("work complete in 20 seconds\nenter timeout duration: ")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("error: negative value for duration")
		return
	}
	go fillChan(jobs)
	go getFromChan(n, jobs, done)

	<-done
}
