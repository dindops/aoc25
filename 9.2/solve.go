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
		return cmp.Compare(j.area, i.area)
	})

	for i, p := range pairsArr {
		if isWithinBoundaries(p, allCoords) {
			fmt.Println(pairsArr[i].area)
			break
		}
	}
}

func isWithinBoundaries(p Pair, c []Coordinates) bool {
	a := c[p.a]
	b := c[p.b]
	x1, x2 := min(a.x, b.x), max(a.x, b.x)
	y1, y2 := min(a.y, b.y), max(a.y, b.y)
	// for each rectangle (pair) test every edge
	// y1 < ey < y2 && x2 > ex1 && x1 < ex2 - horizontal check
	for i := range c {
		ea := c[i]
		eb := c[(i+1)%len(c)]
		ex1, ex2 := min(ea.x, eb.x), max(ea.x, eb.x)
		ey1, ey2 := min(ea.y, eb.y), max(ea.y, eb.y)
		if ey1 == ey2 { //horizontal edge
			if y1 < ey1 && ey1 < y2 && x2 > ex1 && x1 < ex2 {
				return false
			}
		} else { // vertical edge
			if x1 < ex1 && ex1 < x2 && y2 > ey1 && y1 < ey2 {
				return false
			}
		}
	}
	return true
}

func calcArea(a, b Coordinates) int {
	width := max(a.x, b.x) - min(a.x, b.x) + 1
	height := max(a.y, b.y) - min(a.y, b.y) + 1
	area := height * width
	return area
}
