package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Expected: 2401
func part1() {
	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(buf)), "\n")
	grid := make([][]rune, len(lines))
	count := 0

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	target := "XMAS"

	offsetDirections := [][2]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{-1, -1},
		{1, -1},
		{1, 1},
	}

	for rix := range grid {
		for cix := 0; cix < len(grid[rix]); cix++ {
			for _, dir := range offsetDirections {
				offsetRowIx := dir[0]
				offsetColIx := dir[1]
				isTargetFound := true

				for charIndex := 0; charIndex < len(target); charIndex++ {
					rowOffset := rix + (offsetRowIx * charIndex)
					colOffset := cix + (offsetColIx * charIndex)

					if rowOffset < 0 || rowOffset >= len(grid) || colOffset < 0 || colOffset >= len(grid[rix]) {
						isTargetFound = false
						break
					}

					if grid[rowOffset][colOffset] != rune(target[charIndex]) {
						isTargetFound = false
						break
					}
				}

				if isTargetFound {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
