package main

import (
	"aoc2024/guard"
	"fmt"
	"log"
	"os"
	"strings"
)

// Predict guard movements.
// The guard start at a random square facing upward (^ symbol)
//   1. Guard walks forward until an obstacle (# symbol) is hit
//   2. Guard then turn right 90 degrees
// Steps 1 and 2 repeat until guard leaves the map
// Count the number of positions the guard visits along the way.
func main() {

	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	grid := strings.Split(string(input), "\r\n")

	guard := guard.Guard{Direction: guard.DirectionUp, Steps: 1}
	guard.GetStartingPosition(grid)
	fmt.Println(guard.Position)

	for {
		done, err := guard.Move(grid)
		if err != nil {
			fmt.Println(err)
			break
		}
		if done {
			break
		}
	}
	fmt.Printf("Steps: %d\n", guard.Steps)
	fmt.Printf("Visited Positions: %d\n", len(guard.VisitedPositions))
}