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
	var safeReportsWithDampenerCounter int
	for _, report := range reports {
		counter := checkIfSafeWithDampener(report)
		if counter == 0 {
			safeReportsCounter += 1
			safeReportsWithDampenerCounter += 1
		}
		if counter == 1 {
			safeReportsWithDampenerCounter += 1
		}
	}
	fmt.Printf("Safe reports: %d\n", safeReportsCounter)
	fmt.Printf("Safe dampened reports: %d\n", safeReportsWithDampenerCounter)

}

// Report is safe if:
//
// - The levels are either all increasing or all decreasing.
//
// - Any two adjacent levels differ by at least one and at most three.
//
// We tolerate a single bad level in what would otherwise be a safe report.
func checkIfSafeWithDampener(report []int) int {
	var badLevelCounter int
	var lastDigit int
	var increasing bool
	var decreasing bool
	for i, currDigit := range report {

		if i == 0 {
			continue
		}
		lastDigit = report[i-1]

		if currDigit == lastDigit {
			badLevelCounter += 1
		}
		if decreasing && currDigit > lastDigit {
			badLevelCounter += 1
		}
		if increasing && currDigit < lastDigit {
			badLevelCounter += 1
		}

		var diff int
		if currDigit > lastDigit {
			diff = currDigit - lastDigit
		} else {
			diff = lastDigit - currDigit
		}
		if diff > 3 {
			badLevelCounter += 1
		}

		if !decreasing && !increasing && currDigit > lastDigit {
			increasing = true
		}
		if !decreasing && !increasing && currDigit < lastDigit {
			decreasing = true
		}
	}
	return badLevelCounter
}
