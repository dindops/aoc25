package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Coordinates struct {
	x, y int
}

type Pair struct {
	a, b int // indices of coords
	area int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.SplitSeq(strings.TrimSpace(string(data)), "\n")
	var allCoords []Coordinates
	for l := range lines {
		splitLine := strings.SplitSeq(string(l), ",")
		var coords Coordinates
		counter := 0
		for c := range splitLine {
			cInt, err := strconv.Atoi(c)
			if err != nil {
				log.Fatalf("not a number: %s", err)
			}
			switch counter {
			case 0:
				coords.x = cInt
			case 1:
				coords.y = cInt
			}
			counter += 1
		}
		allCoords = append(allCoords, coords)
	}
	var pairsArr []Pair
	n := len(allCoords)
	for i := range n {
		for j := i + 1; j < n; j++ {
			var nPair Pair
			nPair.a = i
			nPair.b = j
			nPair.area = calcArea(allCoords[i], allCoords[j])
			pairsArr = append(pairsArr, nPair)

		}
	}
	slices.SortFunc(pairsArr, func(i, j Pair) int {
		return cmp.Compare(i.area, j.area)
	})
	fmt.Println(pairsArr[len(pairsArr)-1].area)
}

func calcArea(a, b Coordinates) int {
	width := max(a.x, b.x) - min(a.x, b.x) + 1
	height := max(a.y, b.y) - min(a.y, b.y) + 1
	area := height * width
	return area
}
