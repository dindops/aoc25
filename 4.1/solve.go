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
	content := strings.TrimSpace(string(data))
	rows := strings.Split(content, "\n")
	nr := len(rows)
	nc := len(rows[0])
	var directions = [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	noAccessibleRolls := 0
	for r, row := range rows {
		for c, ch := range row {
			if string(ch) == "@" {
				adjacentRolls := 0
				for _, dirs := range directions {
					neighbourR, neighbourC := r+dirs[0], c+dirs[1]
					if neighbourR >= 0 && neighbourR < nr && neighbourC >= 0 && neighbourC < nc {
						if string(rows[neighbourR][neighbourC]) == "@" {
							adjacentRolls += 1
						}
					}
				}
				if adjacentRolls < 4 {
					noAccessibleRolls++
				}
			}
		}
	}
	fmt.Println(noAccessibleRolls)
}
