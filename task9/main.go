package main

import (
	"fmt"
	"sync"
)

func generate(wg *sync.WaitGroup, in chan<- int) {
	for i := 0; i < 51; i++ {
		in <- i
	}
	close(in)
	wg.Done()
}

func makeCalculations(wg *sync.WaitGroup, in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
	close(out)
	wg.Done()
}

func printResult(wg *sync.WaitGroup, out <-chan int) {
	for num := range out {
		fmt.Println(num)
	}
	wg.Done()
}
func main() {
	in := make(chan int)
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go generate(&wg, in)
	go makeCalculations(&wg, in, out)
	go printResult(&wg, out)
	wg.Wait()
}
