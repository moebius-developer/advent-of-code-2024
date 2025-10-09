package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Expected: 1822
func part2() {
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

	offsetDirections := [][2][2]int{
		{
			{-1, 1},
			{1, -1},
		},
		{
			{-1, -1},
			{1, 1},
		},
	}

	for rix := range grid {
		for cix := 0; cix < len(grid[rix]); cix++ {

			if grid[rix][cix] != 'A' {
				continue
			}

			isTargetFound := true

			for _, dir := range offsetDirections {
				offsetRowIx0 := rix + dir[0][0]
				offsetColIx0 := cix + dir[0][1]

				offsetRowIx1 := rix + dir[1][0]
				offsetColIx1 := cix + dir[1][1]

				if offsetRowIx0 < 0 ||
					offsetRowIx0 >= len(grid) ||
					offsetColIx0 < 0 ||
					offsetColIx0 >= len(grid[rix]) ||
					offsetRowIx1 < 0 ||
					offsetRowIx1 >= len(grid) ||
					offsetColIx1 < 0 ||
					offsetColIx1 >= len(grid[rix]) {
					isTargetFound = false
					break
				}

				if (grid[offsetRowIx0][offsetColIx0] == 'M' && grid[offsetRowIx1][offsetColIx1] == 'S') ||
					(grid[offsetRowIx0][offsetColIx0] == 'S' && grid[offsetRowIx1][offsetColIx1] == 'M') {
					continue
				}

				isTargetFound = false
				break
			}

			if isTargetFound {
				count++
			}
		}
	}

	fmt.Println(count)
}
