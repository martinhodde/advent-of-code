package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	registry.Register(6, 1, SolveDay6Part1)
	registry.Register(6, 2, SolveDay6Part2)
}

func SolveDay6Part1(input []string) string {
	allOperands, err := ParseOperands(input)
	if err != nil {
		return fmt.Sprintf("Error parsing operands: %v", err)
	}
	operators := ParseOperators(input)
	operationSum := OperationSum(allOperands, operators)
	return fmt.Sprintf("The total sum of all answers to the individual problems is: %d", operationSum)
}

func SolveDay6Part2(input []string) string {
	return ""
}

// OperationSum computes the total sum of evaluated operations across all operands.
func OperationSum(allOperands [][]int, operators []rune) int {
	total := 0
	for i, operator := range operators {
		total += EvaluateOperation(allOperands, operator, i)
	}

	return total
}

// EvaluateOperation evaluates the operation for the given operator and index across all operands.
func EvaluateOperation(allOperands [][]int, operator rune, index int) int {
	var result int
	switch operator {
	case '+':
		result = 0
		for _, operands := range allOperands {
			result += operands[index]
		}
	case '*':
		result = 1
		for _, operands := range allOperands {
			result *= operands[index]
		}
	}

	return result
}

// ParseOperands parses the operands from the input lines until it encounters an operator line.
func ParseOperands(input []string) ([][]int, error) {
	var allOperands [][]int
	for _, line := range input {
		parts := strings.Fields(line)
		if parts[0][0] == '*' || parts[0][0] == '+' {
			break // We have encountered an operator, so we are done
		}

		var currOperands []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			currOperands = append(currOperands, num)
		}

		allOperands = append(allOperands, currOperands)
	}

	return allOperands, nil
}

// ParseOperators parses the operators from the end of the input lines.
func ParseOperators(input []string) []rune {
	var operators []rune
	for _, line := range input {
		parts := strings.Fields(line)

		// Iterate until we find the first operator line
		if parts[0][0] == '*' || parts[0][0] == '+' {
			for _, part := range parts {
				operators = append(operators, rune(part[0]))
			}
		}
	}

	return operators
}
