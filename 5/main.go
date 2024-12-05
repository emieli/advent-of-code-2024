package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Safety manual updates updates are not printed correctly.
// The new pages must be printed in a specific order.
// Page ordering rules: X|Y must be printed in the correct order. Page X must be printed before page Y.
// Page numbers in each update: X,Y,Z,etc
func main() {

	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	orderingRules := formatOrderingRules(string(input))
	pageUpdates := formatPageUpdates(string(input))

	// Part 1
	var middleNumberCounter int
	var reorderedMiddleNumberCounter int
	for _, pages := range pageUpdates {
		valid := validatePagesInUpdate(pages, orderingRules)
		if !valid {
			reorderedPages := reorderPageUpdate(pages, orderingRules)
			reorderedMiddleNumberCounter += getMiddleNumber(reorderedPages)
			continue
		}
		middleNumberCounter += getMiddleNumber(pages)
	}
	fmt.Println(middleNumberCounter)
	fmt.Println(reorderedMiddleNumberCounter)
}

func getMiddleNumber(numbers []int) int {
	return numbers[len(numbers)/2]
}

// Create a new slice with pages in the same order, but only if
// our page doesn't have any preceeding pages in its
// pagesThatMustComeAfterThisPage rules
func reorderPageUpdate(pages []int, rules map[int][]int) []int {

	newPages := make([]int, 0, len(pages))
	for _, page := range pages {
		pagesThatMustComeAfterThisPage := rules[page]
		var inserted bool
		for nPageIndex, nPage := range newPages {
			if slices.Contains(pagesThatMustComeAfterThisPage, nPage) {
				newPages = slices.Insert(newPages, nPageIndex, page)
				inserted = true
				break
			}
		}
		if !inserted {
			newPages = append(newPages, page)
		}
	}
	return newPages
}

// Go through the pages in reverse. For each page, check the pages before it.
// If any of the pages before our page exist in PagesThatMustComeAfter,
// then page update is invalid.
func validatePagesInUpdate(pages []int, rules map[int][]int) bool {

	for aIndex := len(pages) - 1; aIndex > 0; aIndex-- {
		pageA := pages[aIndex]
		pagesThatMustComeAfterA := rules[pageA]
		for bIndex := 0; bIndex < aIndex; bIndex++ {
			pageB := pages[bIndex]
			if slices.Contains(pagesThatMustComeAfterA, pageB) {
				// page b was found before page a, this is not allowed.
				return false
			}
		}
	}
	return true
}

func formatPageUpdates(input string) [][]int {

	var pageUpdates [][]int

	for _, line := range strings.Split(string(input), "\r\n") {

		var numbers []int

		if !strings.Contains(line, ",") {
			continue
		}
		for _, number := range strings.Split(line, ",") {
			numberAsInt, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println(err)
				continue
			}
			numbers = append(numbers, numberAsInt)
		}

		pageUpdates = append(pageUpdates, numbers)

	}
	return pageUpdates
}

func formatOrderingRules(input string) map[int][]int {

	orderingRules := make(map[int][]int)
	for _, line := range strings.Split(string(input), "\r\n") {

		if !strings.Contains(line, "|") {
			continue
		}

		split := strings.Split(line, "|")

		before, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		after, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}

		orderingRules[before] = append(orderingRules[before], after)
	}
	return orderingRules
}
