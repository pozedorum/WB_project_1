package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool, len(sequence))
	for _, value := range sequence {
		if _, ok := set[value]; !ok {
			set[value] = true
		}
	}

	fmt.Print("{")
	for key := range set {
		fmt.Printf(" \"%s\"", key)
	}
	fmt.Println("}")
}
