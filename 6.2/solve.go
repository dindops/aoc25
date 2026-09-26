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

	var groups [][]string
	for _, r := range columnRanges {
		var cGroup []string
		for _, line := range expandedLines {
			group := line[r.Start:r.End]
			cGroup = append(cGroup, group)
		}
		groups = append(groups, cGroup)
	}
	// Main calculation
	result := 0
	for _, g := range groups {
		operator := strings.TrimSpace(string(g[len(g)-1]))
		var ogMat [][]string
		for i := 0; i < len(g)-1; i++ { // only rows with numbers
			var mRow []string
			for _, n := range g[i] {
				mRow = append(mRow, string(n))
			}
			ogMat = append(ogMat, mRow)
		}
		tMat := transpose(ogMat)
		var collection []string
		for r := range tMat {
			number := strings.TrimSpace(strings.Join(tMat[r], ""))
			collection = append(collection, number)
		}
		cResult := calcCollection(operator, collection)
		result += cResult
	}
	fmt.Println(result)
}

func transpose(grid [][]string) [][]string {
	nRows := len(grid)
	nCols := len(grid[0])

	result := make([][]string, nCols)
	for c := range result {
		result[c] = make([]string, nRows)
	}

	for r := range nRows {
		for c := range nCols {
			result[c][r] = grid[r][c]
		}
	}
	return result
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
