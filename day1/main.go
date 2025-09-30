package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func sortedLeftRight(inputFile string) ([]int, []int) {
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

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	sort.Ints(leftNums)
	sort.Ints(rightNums)

	return leftNums, rightNums
}
