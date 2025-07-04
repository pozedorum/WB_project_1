package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значения квадратов чисел,
взятых из массива [2,4,6,8,10], и выведет результаты в stdout.

Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.
*/

func makeCalculations(input []int) {
	var wg sync.WaitGroup
	for num := range input {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n * n)
		}(num)
		wg.Wait()
	}
}

func main() {
	input := []int{2, 4, 6, 8, 10}
	makeCalculations(input)
}
