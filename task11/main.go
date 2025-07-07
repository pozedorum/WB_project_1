package main

import (
	"fmt"
	"math/rand"
)

func generateMap() map[int]bool {
	newMap := make(map[int]bool, 10)
	for i := 0; i < 10; i++ {
		val := rand.Int() % 10
		if _, ok := newMap[val]; !ok {
			newMap[val] = true
		}
	}
	return newMap
}

func getIntersection(set1, set2 map[int]bool) {
	fmt.Print("{")
	for num := range set1 {
		if _, ok := set2[num]; ok {
			fmt.Printf(" %d", num)
		}
	}
	fmt.Println("}")
}

func main() {
	set1 := generateMap()
	set2 := generateMap()
	getIntersection(set1, set2)
}
