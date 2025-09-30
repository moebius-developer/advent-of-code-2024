package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := getReports(os.Args[1])
	part1(reports)
}

func getReports(inputFile string) [][]int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports [][]int

	for scanner.Scan() {
		line := scanner.Text()
		var rep []int
		for _, nstr := range strings.Fields(line) {
			num, err := strconv.Atoi(nstr)
			if err != nil {
				log.Fatalf("Error converting string to int: %v", err)
			}
			rep = append(rep, num)
		}

		reports = append(reports, rep)
	}

	return reports
}
