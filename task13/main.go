package main

import "fmt"

func main() {
	val1 := 10
	val2 := 25
	val1 = val1 + val2
	val2 = val1 - val2
	val1 = val1 - val2
	fmt.Println(val1, val2)
}

/*
val1
val2

val1 = val1 + val2
val2 = val1 - val2 = val1
val1 = val1 - val2 = val2
*/
