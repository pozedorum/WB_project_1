package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func isUnique(str string) bool {
	str = strings.ToLower(str)
	uniqueMap := make(map[rune]bool, utf8.RuneCount([]byte(str)))

	for _, letter := range str {
		if _, ok := uniqueMap[letter]; !ok {
			uniqueMap[letter] = true
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isUnique("abqwety")) //true
	fmt.Println(isUnique("qqqqq"))   //false
}

//Я воспользовался мапой, так как у неё константный по времени доступ к данным,
//то есть проверка выполняется очень быстро
