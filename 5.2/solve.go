package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(string(data), "\n\n")
	rLines := strings.Split(strings.TrimSpace(parts[0]), "\n")
	ranges := getRanges(rLines)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})
	mergedRanges := mergeRanges(ranges)
	count := 0
	for _, r := range mergedRanges {
		diff := r.end - r.start + 1
		count += diff
	}
	fmt.Println(count)
}

func mergeRanges(ranges []Range) []Range {
	var mergedRanges []Range
	currentRange := ranges[0]
	for _, r := range ranges[1:] {
		if r.start <= currentRange.end {
			currentRange.end = max(currentRange.end, r.end)
		} else {
			mergedRanges = append(mergedRanges, currentRange)
			currentRange = r
		}
	}
	mergedRanges = append(mergedRanges, currentRange)
	return mergedRanges
}

func getRanges(rLines []string) []Range {
	var ranges []Range
	for _, line := range rLines {
		l, r, found := strings.Cut(line, "-")
		if !found {
			log.Fatalf("Unexpected line format: %q", line)
		}
		lInt, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalf("not a number: %s", err)
		}
		rInt, err := strconv.Atoi(r)
		if err != nil {
			log.Fatalf("not a number: %s", err)
		}
		rStruct := Range{lInt, rInt}
		ranges = append(ranges, rStruct)
	}
	return ranges
}
