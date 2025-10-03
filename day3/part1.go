package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Expected: 175615763
func part1() {
	numPairs := getNumberPairs(os.Args[1])

	total := 0
	for _, pair := range numPairs {
		subTotal := 1
		for _, num := range pair {
			subTotal *= num
		}
		total += subTotal
	}

	fmt.Println(total)
}

func getNumberPairs(inputFilename string) [][]int {
	buf, err := os.ReadFile(inputFilename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	re := regexp.MustCompile(`mul\([0-9]+?,[0-9]+?\)`)

	matches := re.FindAllString(string(buf), -1)
	pairs := [][]int{}
	for _, mtch := range matches {
		dstr := strings.Replace(string(mtch), "mul(", "", 1)
		dstr = strings.Replace(dstr, ")", "", 1)
		digitsStr := strings.Split(dstr, ",")

		digits := []int{}
		for _, str := range digitsStr {
			integer, err := strconv.Atoi(str)
			if err != nil {
				log.Fatalf("Failed to convert string to integer: %v", err)
			}

			digits = append(digits, integer)
		}
		if len(digits) != 2 {
			panic("...")
		}

		pairs = append(pairs, digits)
	}

	return pairs
}
