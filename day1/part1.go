package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var leftNums, rightNums []int

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		lnum, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		rnum, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}
		leftNums = append(leftNums, lnum)
		rightNums = append(rightNums, rnum)
	}

	sort.Ints(leftNums)
	sort.Ints(rightNums)

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

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}
