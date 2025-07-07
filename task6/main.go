package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// Все методы кроме первого прерывают работу горутины, позволяя ей выполнить 2 из 4 задач

func process(job int) {
	time.Sleep(time.Second)
	fmt.Printf("job %d is done\n", job)
}

func firstWay() {
	// способ 1: стандартное завершение горутины (например через закрытие канала входных данных)
	// Я бы сюда отнёс и завершения по условию через return,
	jobs := make(chan int, 4)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	jobs <- 4
	done := make(chan struct{})
	go func() {
		defer close(done)
		for job := range jobs {
			process(job)
		}
		fmt.Println("goroutine 1 is ended by closing jobs")

	}()

	close(jobs)
	<-done
}

func secondWay() {
	// способ 2: завершение горутины через дополнительный поток (при его закрытии должна завершаться горутина)
	// Я использую sleep, чтобы поймать момент, когда горутина обработает 2 из 4 процессов
	// и успеет сделать все завершающие процессы после этого
	jobs := make(chan int, 4)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	jobs <- 4
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("goroutine 2 is ended by closing done")
				return
			default:
				job, ok := <-jobs
				if !ok {
					fmt.Println("jobs closed")
					return
				}
				process(job)

			}
		}
	}()

	time.Sleep(2 * time.Second)
	close(done)
	time.Sleep(time.Second + 100*time.Millisecond)
}

func thirdWay() {
	// способ 3: завершение горутины через контекст
	// (аналогичен предыдущему варианту, но работает быстрее и лучше)
	jobs := make(chan int, 4)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	jobs <- 4

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctxt context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine 3 is ended by canceling context")
				return
			default:
				job, ok := <-jobs
				if !ok {
					fmt.Println("jobs closed")
					return
				}
				process(job)

			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func fourthWay() {
	// Способ 4: завершение через time.After, срабатывает после некоторого времени и завершает горутину
	jobs := make(chan int, 4)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	jobs <- 4
	done := make(chan struct{})

	go func() {
		timeout := time.After(2 * time.Second)
		defer close(done)
		for {
			select {
			case <-timeout:
				fmt.Println("goroutine 4 is ended by timeout")
				return
			default:
				job, ok := <-jobs
				if !ok {
					fmt.Println("jobs closed")
					return
				}
				process(job)

			}
		}
	}()

	<-done
}

func fifthWay() {
	// метод 5: использование функции завершения горутины runtime.Goexit()
	jobs := make(chan int, 4)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	jobs <- 4
	done := make(chan struct{})

	go func() {
		defer fmt.Println("goroutine 5 is ended by runtime.Goexit()")
		defer close(done)
		for job := range jobs {
			if job == 3 {
				runtime.Goexit()
			}
			process(job)
		}

	}()

	<-done
}

func main() {
	firstWay()
	secondWay()
	thirdWay()
	fourthWay()
	fifthWay()
}
