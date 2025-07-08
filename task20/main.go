package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	defer out.Flush()
	defer fmt.Fprintln(out)

	in.Scan()
	for _, word := range strings.Split(in.Text(), " ") {
		defer fmt.Fprint(out, word, " ")
	}

}
