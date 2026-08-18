package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	linesMap := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for i := 0; input.Scan(); i++ {
		linesMap[input.Text()] = i

		if _, ok := linesMap["end"]; ok {
			break
		}
	}

	for line, n := range linesMap {

		fmt.Printf("Index - %d\t Value - %s\n", n, line)
	}
}
