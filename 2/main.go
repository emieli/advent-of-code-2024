package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var reports [][]int
	for _, line := range strings.Split(string(input), "\r\n") {

		if line == "" {
			continue
		}

		var ints []int
		for _, characters := range strings.Split(line, " ") {

			digits, err := strconv.Atoi(characters)
			if err != nil {
				fmt.Println(err)
				continue
			}
			ints = append(ints, digits)
		}
		reports = append(reports, ints)
	}

	var safeReportsCounter int
	for _, report := range reports {
		safe := checkIfSafe(report)
		if safe {
			safeReportsCounter += 1
		}
		fmt.Println(safe, report)
	}
	fmt.Printf("Safe reports: %d", safeReportsCounter)

}

// Report is safe if:
//
// - The levels are either all increasing or all decreasing.
//
// - Any two adjacent levels differ by at least one and at most three.
func checkIfSafe(report []int) bool {
	var lastDigit int
	var increasing bool
	var decreasing bool
	for i, currDigit := range report {

		if i == 0 {
			continue
		}

		lastDigit = report[i-1]

		if currDigit == lastDigit {
			fmt.Printf("Duplicate %d %d | ", currDigit, lastDigit)
			return false
		}

		if decreasing && currDigit > lastDigit {
			// Unsafe if decreasing numbers suddenly increase
			fmt.Printf("Decreasing but %d > %d | ", currDigit, lastDigit)
			return false
		}
		if increasing && currDigit < lastDigit {
			// Unsafe if increasing numbers suddenly decrease
			fmt.Printf("Increasing but %d < %d | ", currDigit, lastDigit)
			return false
		}

		var diff int
		if currDigit > lastDigit {
			diff = currDigit - lastDigit
		} else {
			diff = lastDigit - currDigit
		}

		if diff > 3 {
			fmt.Printf("diff: %d | ", diff)
			return false
		}

		if !decreasing && !increasing && currDigit > lastDigit {
			increasing = true
		}
		if !decreasing && !increasing && currDigit < lastDigit {
			decreasing = true
		}
	}
	return true
}
