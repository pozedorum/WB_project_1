package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите строку:")
	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	runeStr := []rune(str)
	for i, j := 0, len(runeStr)-1; i < j; i, j = i+1, j-1 {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}

	fmt.Println(string(runeStr))
}
