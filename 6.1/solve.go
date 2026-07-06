package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	collections := make([][]string, len(strings.Fields(lines[0])))
	for _, l := range lines {
		line := strings.Fields(l)
		for c := range len(line) {
			collections[c] = append(collections[c], line[c])
		}
	}
	sum := 0
	for _, c := range collections {
		operation := c[len(c)-1]
		result := calcCollection(operation, c[:len(c)-1])
		sum += result
	}
	fmt.Println(sum)
}

func calcCollection(operation string, collection []string) int {
	result := 0
	combine := func(a, b int) int { return a + b }
	if operation == "*" {
		combine = func(a, b int) int { return a * b }
		result = 1
	}
	for _, n := range collection {
		nInt, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("not a valid number: %s", err)
		}
		result = combine(result, nInt)

	}
	return result
}
