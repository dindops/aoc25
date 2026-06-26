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
	parts := strings.Split(string(data), "\n\n")
	iLines := strings.Split(strings.TrimSpace(parts[1]), "\n")
	ids := getIds(iLines)
	rLines := strings.Split(strings.TrimSpace(parts[0]), "\n")
	ranges := getRanges(rLines)
	count := 0
	for _, id := range ids {
		if isInRanges(id, ranges) {
			count++
		}
	}
	fmt.Println(count)
}

func isInRanges(id int, ranges [][2]int) bool {
	for _, r := range ranges {
		if id >= r[0] && id <= r[1] {
			return true
		}
	}
	return false
}

func getIds(iLines []string) []int {
	var ids []int
	for _, id := range iLines {
		n, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalf("not a number: %s", err)
		}
		ids = append(ids, n)
	}
	return ids
}

func getRanges(rLines []string) [][2]int {
	var ranges [][2]int
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
		rArr := [2]int{lInt, rInt}
		ranges = append(ranges, rArr)
	}
	return ranges
}
