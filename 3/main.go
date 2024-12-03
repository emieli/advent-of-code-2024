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

	var multiplications int
	sections := strings.Split(string(input), "mul(")

	var doIndex int
	var dontIndex int
	for sectionIndex, section := range sections {

		if strings.Contains(section, "do()") {
			// fmt.Printf("do()    found in %d\n", sectionIndex)
			doIndex = sectionIndex
		}

		if strings.Contains(section, "don't()") {
			// fmt.Printf("don't() found in %d\n", sectionIndex)
			dontIndex = sectionIndex
		}

		output, err := getValidSection(section)
		if err != nil {
			continue
		}

		if dontIndex > doIndex {
			// The latest instruction was a don't(), do don't multiply
			// fmt.Printf("Skipped multiply as %d > %d\n", dontIndex, doIndex)
			continue
		}

		multiplication, err := multiply(output)
		if err != nil {
			continue
		}
		// fmt.Println("Performed multiplication")
		multiplications += multiplication
	}
	fmt.Println(multiplications)
}

func getValidSection(section string) (string, error) {

	var output string
	for _, c := range section {
		character := string(c)

		if character == "0" {
			output += character
		} else if character == "1" {
			output += character
		} else if character == "2" {
			output += character
		} else if character == "3" {
			output += character
		} else if character == "4" {
			output += character
		} else if character == "5" {
			output += character
		} else if character == "6" {
			output += character
		} else if character == "7" {
			output += character
		} else if character == "8" {
			output += character
		} else if character == "9" {
			output += character
		} else if character == "," {
			output += character
		} else if character == ")" {
			break
		} else {
			return "", fmt.Errorf("invalid section: %s", section)
		}

	}

	if !strings.Contains(output, ",") {
		return "", fmt.Errorf("invalid section: %s", section)
	}

	return output, nil

}

func multiply(section string) (int, error) {

	split := strings.Split(section, ",")
	a, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, err
	}
	b, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, err
	}

	return a * b, nil

}
