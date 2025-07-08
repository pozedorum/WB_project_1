package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func getBigNum() *big.Int {
	in := bufio.NewReader(os.Stdin)
	numStr, _ := in.ReadString('\n')
	num := new(big.Int)
	num.SetString(strings.TrimSpace(numStr), 10)
	return num
}

func main() {
	fmt.Println("Введите два числа через enter")
	a := getBigNum()
	b := getBigNum()
	fmt.Println("Sum = ", new(big.Int).Add(a, b))
	fmt.Println("Sub = ", new(big.Int).Sub(a, b))
	fmt.Println("Mul = ", new(big.Int).Mul(a, b))
	if b.Sign() != 0 {
		fmt.Println("Div = ", new(big.Int).Div(a, b))
	} else {
		fmt.Println("Div = n/a")
	}

}
