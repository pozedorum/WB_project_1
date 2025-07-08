package main

import "fmt"

func BinSearch(arr []int, val int) int {

	left, right := 0, len(arr)-1
	if arr[left] == val {
		return left
	} else if arr[right] == val {
		return right
	}
	mid := (left + right) / 2
	for right-left > 1 {
		if val < arr[mid] {
			right = mid
		} else if val > arr[mid] {
			left = mid
		} else {
			return mid
		}
		mid = (left + right) / 2
	}
	return -1
	// бинпоиск не иделальный, но мне хотелось самому по памяти написать (ушло 7 минут))
}

func main() {
	arr := []int{1, 2, 2, 3, 6, 8, 9, 10}
	fmt.Println(BinSearch(arr, 4))
	fmt.Println(BinSearch(arr, 6))
	fmt.Println(BinSearch(arr, 9))
}
