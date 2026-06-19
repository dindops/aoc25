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
	fmt.Println(simulate(instructions))
}

func simulate(instructions []string) int {
	position := 50
	count := 0
	for _, inst := range instructions {
		value, err := strconv.Atoi(inst[1:])
		if err != nil {
			log.Panic(err)
		}

		var crossed int
		if inst[0] == 'L' {
			position, crossed = turnLeft(position, value)
		} else {
			position, crossed = turnRight(position, value)
		}
		count += crossed
	}
	return count
}

func turnRight(position, value int) (int, int) {
	count := 0
	for range value {
		position = (position + 1) % 100
		if position == 0 {
			count++
		}
	}
	return position, count
}

func turnLeft(position, value int) (int, int) {
	count := 0
	for range value {
		position = ((position-1)%100 + 100) % 100
		if position == 0 {
			count++
		}
	}
	return position, count
}
