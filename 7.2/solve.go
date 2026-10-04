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
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	emptyLine := strings.Repeat(".", len(lines[0]))
	timelines := map[int]int{}
	cTimelines := 0
	for iRow, r := range lines {
		if r == emptyLine {
			continue
		}
		buffer := map[int]int{}
		for iCol, c := range r {
			if iRow == 0 && c == 'S' {
				buffer[iCol] = 1
			}
			if c == '^' {
				buffer[iCol-1] += timelines[iCol]
				buffer[iCol+1] += timelines[iCol]
			} else {
				buffer[iCol] += timelines[iCol]
			}
		}
		timelines = buffer
	}
	cTimelines += sumTimelines(timelines)
	fmt.Println(cTimelines)
}

func sumTimelines(t map[int]int) int {
	c := 0
	for _, v := range t {
		c += v
	}
	return c
}
