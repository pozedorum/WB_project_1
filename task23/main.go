package main

import "fmt"

type testStruct struct {
	aboba int
}

// Удаление элемента из слайса по индексу (базовые типы)
func RemoveIndPlace(arr []int, ind int) []int {
	return append(arr[:ind], arr[ind+1:]...)
}

// Удаление элемента из слайса по индексу (ссылочные типы)
func RemoveIndPlacePointer(arr []*testStruct, ind int) []*testStruct {
	arr[ind] = nil
	return append(arr[:ind], arr[ind+1:]...)
}

func main() {

	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("До удаления:", numbers)

	numbers = RemoveIndPlace(numbers, 2)
	fmt.Println("После удаления:", numbers)

	structs := []*testStruct{
		{aboba: 1},
		{aboba: 2},
		{aboba: 3},
		{aboba: 4},
	}
	fmt.Println("\nДо удаления:")
	for _, s := range structs {
		fmt.Printf("%+v ", *s)
	}

	structs = RemoveIndPlacePointer(structs, 1)
	fmt.Println("\nПосле удаления:")
	for _, s := range structs {
		fmt.Printf("%+v ", *s)
	}
}

// Я посморел возмонжые способы проверки утечек и все они были достаточно сложными,
// так что я решил их не реализовывать
