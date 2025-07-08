package main

import (
	"fmt"
	"time"
)

func MySleep(t time.Duration) {
	timer := time.After(t)
	for {
		select {
		case <-timer:
			return
		}
	}
}

func main() {
	var timer int
	fmt.Print("Input how many seconds should program \"sleep\": ")
	fmt.Scan(&timer)
	MySleep(time.Second * time.Duration(timer))
	fmt.Printf("It slept for %d seconds\n", timer)
}
