package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	left, right, err := populateLeftAndRight()
	if err != nil {
		log.Fatal(err)
	}

	sortedLeft := sort(left)
	sortedRight := sort(right)

	// Part One
	distances := make([]int, 1000)
	var totalDistance int
	for i, _ := range sortedLeft {
		distance, err := getDistance(sortedLeft[i], sortedRight[i])
		if err != nil {
			log.Fatal(err)
		}
		distances[i] = distance
		totalDistance += distance
	}
	fmt.Printf("Total distance: %d\n", totalDistance)
}

func populateLeftAndRight() ([]int, []int, error) {

	input, err := os.ReadFile("input")
	if err != nil {
		return nil, nil, err
	}

	left := make([]int, 1000)
	right := make([]int, 1000)

	sep := "   "
	for i, line := range strings.Split(string(input), "\r\n") {
		if line == "" {
			continue
		}
		line_split := strings.Split(line, sep)

		left[i], err = strconv.Atoi(line_split[0])
		if err != nil {
			return nil, nil, err
		}
		right[i], err = strconv.Atoi(line_split[1])
		if err != nil {
			return nil, nil, err
		}
	}

	return left, right, nil
}

func getDistance(a int, b int) (int, error) {
	fmt.Println(a, b)
	if b > a {
		return b - a, nil
	} else if a > b {
		return a - b, nil
	} else if a == b {
		return 0, nil
	} else {
		return 0, fmt.Errorf("couldn't get distance between %d and %d", a, b)
	}
}

func sort(a []int) []int {
	b := make([]int, 0, len(a))
	for _, aInt := range a {
		inserted := false
		for bIndex, bInt := range b {
			if aInt < bInt {
				b = slices.Insert(b, bIndex, aInt)
				inserted = true
				break
			}
		}
		if !inserted {
			b = append(b, aInt)
		}
	}
	return b
}
