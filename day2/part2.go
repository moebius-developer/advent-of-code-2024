package main

import (
	"fmt"
)

func part2(reports [][]int) {
	var safeReports int
	for _, rep := range reports {
		if (isAscending(rep) || isDescending(rep)) && hasNoUnsafeGaps(rep) {
			safeReports++
			continue
		}

		if isSafeWhenDampened(rep) {
			safeReports++
			continue
		}

	}
	fmt.Println(safeReports)
}

func isSafeWhenDampened(rep []int) bool {
	for ix := range len(rep) {
		clone := make([]int, len(rep))
		copy(clone, rep)
		if ix >= len(clone) {
			continue
		}
		clone = append(clone[:ix], clone[ix+1:]...)
		if (isAscending(clone) || isDescending(clone)) && hasNoUnsafeGaps(clone) {
			return true
		}
	}

	return false
}
