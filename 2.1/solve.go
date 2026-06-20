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
	content := strings.TrimSpace(string(data))
	ranges := strings.Split(content, ",")
	sum := 0
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		lo, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		hi, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		sum += iterateOverRange(lo, hi)

	}
	fmt.Println(sum)
}

func iterateOverRange(min, max int) int {
	sum := 0
	for x := min; x <= max; x++ {
		if isInvalid(x) {
			sum += x
		}
	}
	return sum
}

func isInvalid(n int) bool {
	s := strconv.Itoa(n)
	if len(s)%2 != 0 {
		return false
	}
	half := len(s) / 2
	return s[:half] == s[half:]
}
