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
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	var expandedLines []string
	for _, line := range lines {
		if len(line) < maxLen {
			suffix := strings.Repeat(" ", maxLen-len(line))
			expandedLines = append(expandedLines, line+suffix)
		} else {
			expandedLines = append(expandedLines, line)
		}
	}
	separators := make([]int, 0)
	for i := 0; i < maxLen; i++ {
		allBlank := true
		for _, line := range expandedLines {
			if line[i] != ' ' {
				allBlank = false
				break
			}
		}
		if allBlank {
			separators = append(separators, i)
		}
	}

	type ColumnRange struct {
		Start, End int
	}

	var columnRanges []ColumnRange
	start := 0
	for _, sep := range separators {
		var cRange ColumnRange
		cRange.Start = start
		cRange.End = sep
		columnRanges = append(columnRanges, cRange)
		start = sep + 1
	}
	var lastRange ColumnRange
	lastRange.Start = separators[len(separators)-1] + 1
	lastRange.End = maxLen
	columnRanges = append(columnRanges, lastRange)

	numberRows := expandedLines[:len(expandedLines)-1]
	operatorRow := expandedLines[len(expandedLines)-1]
	result := 0

	for _, r := range columnRanges {
		operator := strings.TrimSpace(operatorRow[r.Start:r.End])
		var collection []string
		for c := r.Start; c < r.End; c++ {
			sb := strings.Builder{}
			for _, row := range numberRows {
				sb.WriteByte(row[c])
			}
			collection = append(collection, strings.TrimSpace(sb.String()))
		}
		result += calcCollection(operator, collection)
	}
	fmt.Println(result)
}

func calcCollection(operator string, collection []string) int {
	result := 0
	combine := func(a, b int) int { return a + b }
	if operator == "*" {
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
