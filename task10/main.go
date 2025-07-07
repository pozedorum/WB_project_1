package main

import "fmt"

func getGroupKey(temp float32) int {
	return int(temp) / 10 * 10
}

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float32)

	for _, temp := range temperatures {
		key := getGroupKey(temp)
		groups[key] = append(groups[key], temp)
	}
	for key, value := range groups {
		fmt.Printf("%d: {", key)
		for _, val := range value {
			fmt.Printf(" %.1f", val)
		}
		fmt.Println("}")
	}
}
