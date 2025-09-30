package main

import (
	"fmt"
	"math"
)

func part1(reports [][]int) {
	var safeReports int
	for _, rep := range reports {
		if (isAscending(rep) || isDescending(rep)) && hasNoUnsafeGaps(rep) {
			safeReports++
			continue
		}
	}
	fmt.Println(safeReports)
}

func isAscending(rep []int) bool {
	return _loop(rep, func(prev, num int) bool {
		return prev <= num
	})
}

func isDescending(rep []int) bool {
	return _loop(rep, func(prev, num int) bool {
		return prev >= num
	})
}

func hasNoUnsafeGaps(rep []int) bool {
	return _loop(rep, func(prev, num int) bool {
		abs := math.Abs(float64(num - prev))
		return !(abs < 1 || abs > 3)
	})
}

func _loop(rep []int, logic func(prev, num int) bool) bool {
	prev := rep[0]
	for _, num := range rep[1:] {
		if !logic(prev, num) {
			return false
		}
		prev = num
	}
	return true
}
