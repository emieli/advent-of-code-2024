package guard

import (
	"fmt"
	"slices"
	"strings"
)

type Direction string
const (
	DirectionUp Direction = "up"
	DirectionDown Direction = "down"
	DirectionLeft Direction = "left"
	DirectionRight Direction = "right"
)

type Position struct {
	y int
	x int
}

type PositionState string
const (
	PositionStateClear PositionState = "clear"
	PositionStateObstacle PositionState = "obstacle"
	PositionStateOob PositionState = "out of bounds"
)

type Guard struct {
	Direction Direction
	Position Position
	VisitedPositions []Position
	Steps int
}

func (guard *Guard) addVisitedPosition() {
	if slices.Contains(guard.VisitedPositions, guard.Position) {
		fmt.Println("Did not add duplicate position")
		return
	}
	guard.VisitedPositions = append(guard.VisitedPositions, guard.Position)
	fmt.Println("Added position as visited")
}

func (guard *Guard) Move(grid []string) (bool, error) {

	newPosition := guard.Position
	if guard.Direction == DirectionUp {
		newPosition.y -= 1
	} else if guard.Direction == DirectionDown {
		newPosition.y += 1
	} else if guard.Direction == DirectionLeft {
		newPosition.x -= 1
	} else if guard.Direction == DirectionRight {
		newPosition.x += 1
	} else {
		return false, fmt.Errorf("invalid direction: %s", guard.Direction)
	}

	state, err := getPositionState(newPosition, grid)
	if err != nil {
		return false, err
	}

	fmt.Printf("%d %s\n", newPosition, state)

	if state == PositionStateClear {
		fmt.Printf("Guard moved %s from %d to %d\n", guard.Direction, guard.Position, newPosition)
		guard.Position = newPosition
		guard.addVisitedPosition()
		guard.Steps += 1
		return false, nil
	}
	if state == PositionStateObstacle {
		guard.Turn()
		return false, nil
	}
	if state == PositionStateOob {
		// Guard has left the map
		return true, nil
	}

	return false, fmt.Errorf("Don't know what to do now: %s\n", state)
}

func (guard *Guard) Turn() {
	directionChange := map[Direction]Direction{
		DirectionUp:    DirectionRight,
		DirectionRight: DirectionDown,
		DirectionDown:  DirectionLeft,
		DirectionLeft:  DirectionUp,
	}
	newDirection := directionChange[guard.Direction]
	fmt.Printf("Guard changed direction from %s to %s\n", guard.Direction, newDirection)
	guard.Direction = newDirection
}

func (guard *Guard) GetStartingPosition(grid []string) {

	for lineNumber, line := range grid {
		if !strings.Contains(line, "^") {
			continue
		}
		guard.Position.y = lineNumber
		for charNumber, char := range line {
			if string(char) != "^" {
				continue
			}
			guard.Position.x = charNumber
			return
		}
	}
}

func getPositionState(position Position, grid []string) (PositionState, error) {

	if position.x < 0 {
		// Left map on the left side
		return PositionStateOob, nil
	}
	if position.y < 0 {
		// Left map on the top side
		return PositionStateOob, nil
	}

	if position.x > len(grid[0])-1 {
		// Left map on the right side
		return PositionStateOob, nil
	}
	if position.y > len(grid)-1 {
		// Left map on the bottom side
		return PositionStateOob, nil
	}

	gridPositionSymbol := string(grid[position.y][position.x])

	if gridPositionSymbol == "." {
		return PositionStateClear, nil
	}
	if gridPositionSymbol == "^" {
		return PositionStateClear, nil
	}

	if gridPositionSymbol == "#" {
		return PositionStateObstacle, nil
	}

	return PositionStateOob, fmt.Errorf("Unexpected position symbol: %s\n", gridPositionSymbol)
}