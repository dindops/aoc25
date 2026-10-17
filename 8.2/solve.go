package main

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	// "sort"
	"strconv"
	"strings"
)

type Coordinates struct {
	x, y, z int
}

type Pair struct {
	a, b int // indices of coords
	dist float64
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
			case 2:
				coords.z = cInt
			}
			counter += 1
		}
		allCoords = append(allCoords, coords)
	}
	// loop over coords, create pairs and calculate the distance
	var pairsArr []Pair
	n := len(allCoords)
	for i := range n {
		for j := i + 1; j < n; j++ {
			var nPair Pair
			nPair.a = i
			nPair.b = j
			nPair.dist = calcDistance(allCoords[i], allCoords[j])
			pairsArr = append(pairsArr, nPair)

		}
	}
	// now sort them
	slices.SortFunc(pairsArr, func(i, j Pair) int {
		return cmp.Compare(i.dist, j.dist)
	})

	circuits := make([]int, n)
	for i := range n {
		circuits[i] = i
	}
	var lastPair Pair
	for _, p := range pairsArr {
		if circuits[p.a] != circuits[p.b] {
			lastPair = p
			oldB := circuits[p.b]
			for i, c := range circuits {
				if c == oldB {
					circuits[i] = circuits[p.a]
				}
			}
		}
	}
	res := allCoords[lastPair.a].x * allCoords[lastPair.b].x
	fmt.Println(res)
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func calcDistance(a, b Coordinates) float64 {
	// sqrt of (x1 - x2)^2 + (y1 - y2)^2 + (z1 - z2)^2
	xd := a.x - b.x
	yd := a.y - b.y
	zd := a.z - b.z
	sum := float64(powInt(xd, 2) + powInt(yd, 2) + powInt(zd, 2))
	dist := math.Sqrt(sum)
	return dist
}
