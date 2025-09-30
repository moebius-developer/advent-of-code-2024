package main

import "fmt"

func part2(leftNums, rightNums []int) {
	rightCountByNum := map[int]int{}
	for _, num := range rightNums {
		if _, ok := rightCountByNum[num]; ok {
			rightCountByNum[num]++
			continue
		}

		rightCountByNum[num] = 1
	}

	var simScore int
	for _, num := range leftNums {
		if count, ok := rightCountByNum[num]; ok {
			simScore += num * count
		}
	}
	fmt.Println(simScore)
}
