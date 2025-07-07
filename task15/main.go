package main

import "strings"

/*
var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}


Проблема в том, что при взятии среза, берётся ссылка на ту же строку и после завершения функции
сборщик мусора не может освободить память из-за оставшейся ссылки в justString
В качестве решения надо просто скопировать срез строки в justString, тогда изначальная строка
перестанет иметь указателей и уничтожится по завершению функции
*/

func createHugeString(length int) string {
	return strings.Repeat("a", length)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}
