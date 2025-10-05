package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Expected: 74361272
func part2() {
	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

	matches := re.FindAllStringSubmatch(string(buf), -1)
	do := true
	total := 0
	for _, match := range matches {
		if match[0] == "do()" {
			do = true
			continue
		}
		if match[0] == "don't()" {
			do = false
			continue
		}
		if do {
			pair := extractNumPair(match[1:])
			subTotal := 1
			for _, num := range pair {
				subTotal *= num
			}
			total += subTotal
		}
	}
	fmt.Printf("%#v\n", total)
}

func extractNumPair(strs []string) []int {
	digits := []int{}
	for _, str := range strs {
		integer, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Failed to convert string to integer: %v", err)
		}

		digits = append(digits, integer)
	}
	if len(digits) != 2 {
		panic("...")
	}
	return digits
}
