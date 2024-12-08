package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	Result  int
	Numbers []int
}

// Determine the math operator by the result. Examples:
//
//	190: 10 19 # 10 * 19
//	3267: 81 40 27 # 81 + 40 * 27
//	161011: 16 10 13 # Invalid
//
// Operators are read left to right, not according to predecence rules
// Not all equations are possible.
// Print the total calibration amount, which is the sum of the test values
// produced by valid equations
func main() {

	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	equations := getEquations(string(input))

	var total int
	for _, e := range equations {
		combinations := generateOperatorCombinations(e)
		for _, combination := range combinations {
			if validEquationOperator(e.Numbers, e.Result, combination) {
				total += e.Result
				break
			}
		}
	}
	fmt.Println(total)
}

// A three-number equation require two operators. In that case, we have a total of
// four possible operator combinations: ++, +*, *+, **.
// We use binary to generate all operator combinations, replacing 0 with + and 1 with *.
func generateOperatorCombinations(e Equation) []string {

	totalOperators := len(e.Numbers) - 1
	possibleCombinations := int(math.Pow(2, float64(totalOperators)))
	operatorCombinations := make([]string, 0, possibleCombinations)

	for c := 0; c < possibleCombinations; c++ {

		cToBinaryToString := string(strconv.FormatInt(int64(c), 2))
		// Pad string to make it into a "binary" format
		for len(cToBinaryToString) < totalOperators {
			cToBinaryToString = "0" + cToBinaryToString
		}
		operators := strings.ReplaceAll(cToBinaryToString, "1", "*")
		operators = strings.ReplaceAll(operators, "0", "+")
		operatorCombinations = append(operatorCombinations, operators)
	}
	return operatorCombinations
}

func validEquationOperator(numbers []int, expectedResult int, operators string) bool {

	var total int
	for i := 0; i < len(numbers)-1; i++ {

		if i == 0 {
			total = numbers[i]
		}
		operator := string(operators[i])
		if operator == "+" {
			// fmt.Printf("%d + %d\n", total, numbers[i+1])
			total += numbers[i+1]
		} else if operator == "*" {
			// fmt.Printf("%d x %d\n", total, numbers[i+1])
			total *= numbers[i+1]
		}
	}
	return total == expectedResult
}

func getEquations(input string) []Equation {

	inputLines := strings.Split(string(input), "\r\n")
	equations := make([]Equation, 0, len(inputLines))

	for _, line := range inputLines {
		lineSplitByColon := strings.Split(line, ": ")
		result, err := strconv.Atoi(lineSplitByColon[0])
		if err != nil {
			log.Fatal(err)
		}
		numbers := lineSplitByColon[1]

		equation := Equation{Result: result, Numbers: make([]int, 0, len(numbers))}
		for _, number := range strings.Split(numbers, " ") {
			numberAsInt, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			equation.Numbers = append(equation.Numbers, numberAsInt)
		}
		equations = append(equations, equation)
	}
	return equations

}
