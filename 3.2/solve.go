package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := strings.TrimSpace(string(data))
	banks := strings.Split(content, "\n")
	sum := 0
	for _, b := range banks {
		sum += findHighestJoltage(b)
	}
	fmt.Println(sum)
}

func findHighestJoltage(bank string) int {
	stack := make([]int, 0, len(bank))
	r := len(bank) - 12
	for _, b := range bank {
		n := int(b - '0') // b is of type rune, '0' is rune literal, apparently subtraction gives int
		for r > 0 && len(stack) > 0 && stack[len(stack)-1] < n {
			stack = stack[:len(stack)-1] // pop
			r -= 1
		}
		stack = append(stack, n)
	}
	stack = stack[:12]
	result := 0
	for _, d := range stack {
		result = result*10 + d
	}
	return result
}
