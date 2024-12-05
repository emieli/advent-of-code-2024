package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Find the occurence of XMAS in the input.
// It can be written horizontal, vertical, diagonal, written backwards, or even overlapping other words.
// Print how many times XMAS appear.
func main() {

	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	count := getHorizontalForwardCount(string(input))
	count += getHorizontalBackwardCount(string(input))
	count += getVerticalCount(string(input))
	count += getDiagonalLeftToRightCount(string(input))
	count += getDiagonalRightToLeftCount(string(input))
	fmt.Println(count)

}

func getHorizontalForwardCount(input string) int {
	input = strings.Replace(input, "\r\n", "", -1)
	return len(strings.Split(input, "XMAS"))
}

func getHorizontalBackwardCount(input string) int {
	input = strings.Replace(input, "\r\n", "", -1)
	return len(strings.Split(input, "SAMX"))
}

func getVerticalCount(input string) int {

	lines := strings.Split(input, "\r\n")
	lineLength := len(lines[0])

	var wordCounter int

	for lineIndex := 0; lineIndex < len(lines)-4; lineIndex++ {
		for charIndex := 0; charIndex < lineLength; charIndex++ {
			var word string
			word += string(lines[lineIndex][charIndex])
			word += string(lines[lineIndex+1][charIndex])
			word += string(lines[lineIndex+2][charIndex])
			word += string(lines[lineIndex+3][charIndex])

			if word == "XMAS" {
				wordCounter += 1
			}
			if word == "SAMX" {
				wordCounter += 1
			}
		}
	}
	return wordCounter

}

func getDiagonalLeftToRightCount(input string) int {

	lines := strings.Split(input, "\r\n")
	lineLength := len(lines[0])

	var wordCounter int

	for lineIndex := 0; lineIndex < len(lines)-4; lineIndex++ {
		for charIndex := 0; charIndex < lineLength-4; charIndex++ {
			var word string
			word += string(lines[lineIndex][charIndex])
			word += string(lines[lineIndex+1][charIndex+1])
			word += string(lines[lineIndex+2][charIndex+2])
			word += string(lines[lineIndex+3][charIndex+3])

			if word == "XMAS" {
				wordCounter += 1
			}
			if word == "SAMX" {
				wordCounter += 1
			}
		}
	}
	return wordCounter
}

func getDiagonalRightToLeftCount(input string) int {

	lines := strings.Split(input, "\r\n")
	lineLength := len(lines[0])

	var wordCounter int

	for lineIndex := 0; lineIndex < len(lines)-4; lineIndex++ {
		for charIndex := 3; charIndex < lineLength; charIndex++ {
			var word string
			word += string(lines[lineIndex][charIndex])
			word += string(lines[lineIndex+1][charIndex-1])
			word += string(lines[lineIndex+2][charIndex-2])
			word += string(lines[lineIndex+3][charIndex-3])

			if word == "XMAS" {
				wordCounter += 1
			}
			if word == "SAMX" {
				wordCounter += 1
			}
		}
	}
	return wordCounter
}
