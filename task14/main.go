package main

import "fmt"

func main() {

	intList := []interface{}{"aaa", 123, true, make(chan int)}

	for _, inter := range intList {
		switch v := inter.(type) {
		case int:
			fmt.Printf("Это int: %d\n", v)
		case string:
			fmt.Printf("Это string: %s\n", v)
		case bool:
			fmt.Printf("Это bool: %v\n", v)
		case chan int:
			fmt.Printf("Это chan int: %v\n", v)
		default:
			fmt.Printf("Неизвестный тип: %T\n", v)
		}

	}
}
