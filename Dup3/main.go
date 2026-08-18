package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputCount := make(map[string]int)

	for _, filename := range os.Args[1:] {

		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		words := strings.Fields(string(data))
		index := 0

		for _, word := range words {
			inputCount[string(word)] = index
			index++
		}
	}

	for key, val := range inputCount {
		fmt.Printf("Key - %s: Value - %d\n", key, val)
	}

}
