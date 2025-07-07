package main

import "fmt"

func SetBit(num *int64, place int, bitValue int) error {
	if place < 1 || place > 64 {
		return fmt.Errorf("error: place not in range")
	} else if bitValue != 0 && bitValue != 1 {
		return fmt.Errorf("error: bitValue not in range")
	}
	place--
	if bitValue == 1 {
		*num = *num | int64(1<<place)
	} else {
		*num = *num & ^int64(1<<place)
	}
	return nil
}

func main() {
	var (
		num      int64
		place    int
		bitValue int
	)
	fmt.Print("Entrer number: ")
	fmt.Scan(&num)
	fmt.Print("Entrer bit place to change (1-64): ")
	fmt.Scan(&place)
	fmt.Print("Entrer value of new bit: ")
	fmt.Scan(&bitValue)

	if err := SetBit(&num, place, bitValue); err != nil {
		panic(err)
	} else {
		fmt.Printf("result value: %d\n", num)
		fmt.Printf("binary result value: %b\n", uint64(num))
	}
}
