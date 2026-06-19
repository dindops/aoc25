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
	instructions := strings.Split(string(data), "\n")
	instructions = instructions[:len(instructions)-1]
	fmt.Printf("%d", simulate(instructions))
}

func simulate(instructions []string) int {
	position := 50
	count := 0
	for _, inst := range instructions {
		direction := string(inst[0])
		value, err := strconv.Atoi(inst[1:])
		if err != nil {
			log.Panic(err)
		}

		if direction == "L" {
			position = ((position-value)%100 + 100) % 100 // Go's modulo doesn't handle negative numbers well
		} else {
			position = (position + value) % 100
		}
		if position == 0 {
			count++
		}
		fmt.Printf("%s, %d, %d\n", direction, value, position)
	}
	return count
}
