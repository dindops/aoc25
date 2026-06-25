package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := strings.TrimSpace(string(data))
	rows := strings.Split(content, "\n")
	nr := len(rows)
	nc := len(rows[0])
	noAccessibleRolls := 0
	for r, row := range rows {
		for c, ch := range row {
			if string(ch) == "@" {
				if nr-1 > r && r > 0 && nc-1 > c && c > 0 { // not edge - all neighbours available
					adjPos := rows[r-1][c-1:c+2] + string(rows[r][c-1]) + string(rows[r][c+1]) + rows[r+1][c-1:c+2]
					if isRollAccessible(adjPos) {
						noAccessibleRolls += 1
					}
					continue
				}
				if r == 0 {
					if c == 0 { //top left corner
						adjPos := string(rows[r][c+1]) + rows[r+1][:c+2]
						if isRollAccessible(adjPos) {
							noAccessibleRolls += 1
						}
						continue
					}
					if c == nc-1 { //top right corner
						adjPos := string(rows[r][c-1]) + rows[r+1][:c+1]
						if isRollAccessible(adjPos) {
							noAccessibleRolls += 1
						}
						continue
					}
					if c > 0 && c < nc-1 {
						adjPos := string(rows[r][c-1]) + string(rows[r][c+1]) + rows[r+1][c-1:c+2]
						if isRollAccessible(adjPos) {
							noAccessibleRolls += 1
						}
						continue
					}
				}
				if r == nr-1 {
					if c == 0 { //bottom left corner
						adjPos := string(rows[r][c+1]) + rows[r-1][:c+2]
						if isRollAccessible(adjPos) {
							noAccessibleRolls += 1
						}
						continue
					}
					if c == nc-1 { //bottom right corner
						adjPos := string(rows[r][c-1]) + rows[r-1][:c+1]
						if isRollAccessible(adjPos) {
							noAccessibleRolls += 1
						}
						continue
					}
					if c > 0 && c < nc-1 {
						adjPos := string(rows[r][c-1]) + string(rows[r][c+1]) + rows[r-1][c-1:c+2]
						if isRollAccessible(adjPos) {
							noAccessibleRolls += 1
						}
						continue
					}
				}
				if c == 0 {
					adjPos := string(rows[r+1][:c+2]) + string(rows[r][c+1]) + rows[r-1][:c+2]
					if isRollAccessible(adjPos) {
						noAccessibleRolls += 1
					}
					continue
				}
				if c == nc-1 {
					adjPos := string(rows[r+1][c-1:]) + string(rows[r][c-1]) + rows[r-1][c-1:]
					if isRollAccessible(adjPos) {
						noAccessibleRolls += 1
					}
					continue
				}
			}
		}
	}
	fmt.Println(noAccessibleRolls)
}

func isRollAccessible(adjPos string) bool {
	adjacentRolls := 0
	for _, ach := range adjPos {
		if string(ach) == "@" {
			adjacentRolls += 1
		}
	}
	if adjacentRolls < 4 {
		return true
	}
	return false
}

