package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func init() {
	registry.Register(6, 1, SolveDay6Part1)
	registry.Register(6, 2, SolveDay6Part2)
}

func SolveDay6Part1(input []string) string {
	expressionSum := expressionSum(parseOperands(input), parseOperators(input))
	return fmt.Sprintf("The total sum of all regular math answers is: %d", expressionSum)
}

func SolveDay6Part2(input []string) string {
	expressionSum := cephalopodExpressionSum(input)
	return fmt.Sprintf("The total sum of all cephalopod math answers is: %d", expressionSum)
}

// expressionSum computes the total sum of all evaluated expressions.
func expressionSum(allOperands [][]int, operators []rune) int {
	total := 0
	var operands []int

	for i, operator := range operators {
		for _, row := range allOperands {
			operands = append(operands, row[i])
		}

		total += evaluateExpression(operands, operator)
		operands = operands[:0]
	}

	return total
}

// cephalopodExpressionSum computes the total sum of all evaluated expressions, but assumes
// the input is formatted right to left with operand digits being arranged vertically.
func cephalopodExpressionSum(input []string) int {
	total := 0
	var currOperands []int

	// For cephalopod expressions, it will be easier to operate on the raw input grid instead
	// of pre-fetching parsed integer operands due to the vertical alignment of the digits
	operatorRow := getOperatorRowIndex(input)

	// We use left-aligned operators as our signal to terminate operand accumulation for
	// a given expression, so we iterate through the input from right to left, bottom to top
	// to construct operands
	for j := len(input[0]) - 1; j >= 0; j-- {
		operand, powerOfTen := 0, 0
		for i := operatorRow - 1; i >= 0; i-- {
			operandChar := input[i][j]
			if operandChar >= '0' && operandChar <= '9' {
				digit, _ := strconv.Atoi(string(operandChar))
				operand += digit * int(math.Pow10(powerOfTen))
				powerOfTen++ // Move to the next digit place
			}
		}
		currOperands = append(currOperands, operand)

		// If we encounter a left-aligned operator, we can evaluate the aggregated expression
		operatorChar := input[operatorRow][j]
		if operatorChar == '*' || operatorChar == '+' {
			total += evaluateExpression(currOperands, rune(operatorChar))
			currOperands = currOperands[:0]
			j-- // Skip column of spaces between operators
		}
	}

	return total
}

// getOperatorRowIndex returns the index of the row that contains the operators.
func getOperatorRowIndex(input []string) int {
	for i, line := range input {
		if line[0] == '*' || line[0] == '+' {
			return i
		}
	}

	return -1
}

// evaluateExpression evaluates the expression for the given operator and index across all operands.
func evaluateExpression(operands []int, operator rune) int {
	var result int
	switch operator {
	case '+':
		result = 0
		for _, operand := range operands {
			result += operand
		}
	case '*':
		result = 1
		for _, operand := range operands {
			result *= operand
		}
	}

	return result
}

// parseOperands parses the operand rows from the input lines until it encounters an operator line.
func parseOperands(input []string) [][]int {
	var allOperands [][]int
	for _, line := range input {
		parts := strings.Fields(line)
		if parts[0][0] == '*' || parts[0][0] == '+' {
			break // We have encountered an operator, so we are done
		}

		var currOperands []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			currOperands = append(currOperands, num)
		}
		allOperands = append(allOperands, currOperands)
	}

	return allOperands
}

// parseOperators parses the operators from the end of the input lines.
func parseOperators(input []string) []rune {
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
