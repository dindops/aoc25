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
		h1, idx := findHighestJoltage(b)
		if idx+1 == len(b) {
			h2, _ := findHighestJoltage(b[:idx])
			result := h2*10 + h1
			sum += result
		} else {
			h2, _ := findHighestJoltage(b[idx+1:])
			result := h1*10 + h2
			sum += result
		}
	}
	fmt.Println(sum)
}

func findHighestJoltage(bank string) (int, int) {
	h := 0
	idx := 0
	for i, b := range bank {
		n := int(b - '0') // b is of type rune, '0' is rune literal, apparently subtraction gives int
		if n > h {
			h = n
			idx = i
		}
	}
	return h, idx
}
