package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

// Antennas are tuned to a frequency indicated by a single
// lowercase/uppercase letter or digit.
// The signal only applies its effect at specific antinodes based on the
// resonant frequencies of the antennas.
// An antinode occurs at any point that is perfectly in line with two antennas of the same
// frequency, but only when one antenna is twice as far away as the other.
// So for each two antennas, there are two antinodes, one for either side of them:
//
//	# a a #
//
// Adding a third antenna with the same frequency creates several more antinodes:
//
//	    #
//	#
//	     a        #
//	          a
//	      a        #
//	  #
//	       #
//
// How many unique locations within the map contain antinodes?

type Position struct {
	y, x int
}

type Antenna struct {
	Type     string
	Position Position
}

func getAntennas() map[string][]Antenna {

	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	antennas := make(map[string][]Antenna)
	for y, line := range strings.Split(string(input), "\r\n") {
		for x, c := range line {
			char := string(c)
			if char == "." {
				continue
			}
			antenna := Antenna{Type: char, Position: Position{y: y, x: x}}
			antennas[char] = append(antennas[char], antenna)
		}
	}
	return antennas
}

func getAllAntiNodes(A Antenna, B Antenna) []Position {

	var antiNodes []Position
	antiNodes = append(antiNodes, A.Position, B.Position)

	// If A is on row 4 and B is on row 6, antiNode is placed on row 2.
	// 4 - 6 = -2, and 4 + -2 = 2
	// If A is on row 6 and B is on row 4, antinode is placed on row 8.
	// 6-4 = 2, and 6 + 2 = 8
	yDistance := A.Position.y - B.Position.y
	xDistance := A.Position.x - B.Position.x

	// Find all repeating antinodes to the end of the map
	for {

		var antiNode Position

		antiNode.y = A.Position.y + yDistance
		if antiNode.y < 0 || antiNode.y > 49 {
			// Antinode is out of bounds
			break
		}

		antiNode.x = A.Position.x + xDistance
		if antiNode.x < 0 || antiNode.x > 49 {
			// Antinode is out of bounds
			break
		}

		antiNodes = append(antiNodes, antiNode)

		A.Position.y = antiNode.y
		A.Position.x = antiNode.x

	}
	fmt.Println(A.Position, B.Position, antiNodes)
	return antiNodes
}

func main() {

	var antiNodes []Position

	antennas := getAntennas()
	for atype, antennaType := range antennas {
		fmt.Printf("Processing %s antennas...\n", atype)
		for _, A := range antennaType {
			for _, B := range antennaType {

				if A.Position == B.Position {
					// We're looking at the same antenna
					continue
				}

				// This will add duplicates
				nodes := getAllAntiNodes(A, B)
				if len(nodes) > 0 {
					antiNodes = append(antiNodes, nodes...)
				}
			}
		}
		fmt.Println(len(antiNodes))
	}

	var nonDuplicateAntiNodes []Position
	for _, antinode := range antiNodes {
		if !slices.Contains(nonDuplicateAntiNodes, antinode) {
			nonDuplicateAntiNodes = append(nonDuplicateAntiNodes, antinode)
		}
	}
	fmt.Println(len(nonDuplicateAntiNodes))

}
