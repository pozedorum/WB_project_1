package main

import "fmt"

func quickSort(arr []int) []int {
	return quickSortReq(arr, 0, len(arr)-1)
}

func partitioning(arr []int, left, right int) ([]int, int) {
	pivot := arr[right]

	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[right], arr[i] = arr[i], arr[right]
	return arr, i
	// теперь pivot на своём месте, осталось отсортировать левую и правую часть
}

func quickSortReq(arr []int, left, right int) []int {
	var pivot int
	if left < right {
		arr, pivot = partitioning(arr, left, right) //нашли центральный элемент и отсортировали список относительно него
		arr = quickSortReq(arr, left, pivot-1)      // теперь аналогично сортируем левую и правую части
		arr = quickSortReq(arr, pivot+1, right)
	}
	return arr
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	arr = quickSort(arr)
	for _, i := range arr {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
