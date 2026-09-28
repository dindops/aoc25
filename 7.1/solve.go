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
	beamsIdx := map[int]bool{}
	splitCount := 0
	for iRow, r := range lines {
		if r == emptyLine {
			continue
		}
		for iCol, c := range r {
			if iRow == 0 && c == 'S' {
				beamsIdx[iCol] = true
			}
			if c == '^' {
				if beamsIdx[iCol] == true {
					splitCount += 1
					beamsIdx[iCol] = false
					beamsIdx[iCol-1] = true
					beamsIdx[iCol+1] = true
				}
			}
		}
	}
	fmt.Println(splitCount)
}
