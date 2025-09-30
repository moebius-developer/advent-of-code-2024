package main

import (
	"fmt"
)

func part1(leftNums, rightNums []int) {
	dist := 0
	for ix := range leftNums {
		var diff int
		left := leftNums[ix]
		right := rightNums[ix]
		if left > right {
			diff = left - right
		} else {
			diff = right - left
		}
		dist += diff
	}
	fmt.Println(dist)
}
