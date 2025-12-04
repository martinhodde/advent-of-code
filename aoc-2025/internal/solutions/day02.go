package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	registry.Register(2, 1, SolveDay2Part1)
	registry.Register(2, 2, SolveDay2Part2)
}

func SolveDay2Part1(input []string) string {
	invalidIDSum := CalculateInvalidIDSum(input[0])
	return fmt.Sprintf("The sum of all the invalid IDs is %d", invalidIDSum)
}

func SolveDay2Part2(input []string) string {
	return ""
}

// CalculateInvalidIDSum takes a string representing ID ranges and
// returns the sum of all invalid IDs within those ranges.
func CalculateInvalidIDSum(ranges string) int {
	idRanges, err := ParseIDRanges(ranges)
	if err != nil {
		fmt.Printf("Error parsing ID ranges: %v\n", err)
		return 0
	}

	sum := 0
	for _, id := range FilterInvalidIDs(idRanges) {
		sum += id
	}

	return sum
}

// FilterInvalidIDs takes a slice of [2]int representing ID ranges and
// returns a slice of all invalid IDs within those ranges.
func FilterInvalidIDs(ranges [][2]int) []int {
	var invalidIDs []int
	for _, r := range ranges {
		lower, upper := r[0], r[1]
		for id := lower; id <= upper; id++ {
			if IsInvalidID(id) {
				invalidIDs = append(invalidIDs, id)
			}
		}
	}

	return invalidIDs
}

// IsInvalidID checks if a given ID is invalid based on the criteria
// of being made only of some sequence of digits repeated twice.
func IsInvalidID(id int) bool {
	idString := strconv.Itoa(id)
	numDigits := len(idString)
	return numDigits%2 == 0 && idString[:numDigits/2] == idString[numDigits/2:]
}

// ParseIDRanges takes a line containing ID ranges in the format "1-3,5-7,10-15"
// and returns a slice of [2]int representing the lower and upper bounds of each range.
func ParseIDRanges(line string) ([][2]int, error) {
	inputRanges := strings.Split(line, ",")
	outputRanges := make([][2]int, len(inputRanges))

	for i, r := range inputRanges {
		bounds := strings.Split(r, "-")
		if len(bounds) != 2 {
			return nil, fmt.Errorf("invalid range: %s", r)
		}

		lower, err := strconv.Atoi(bounds[0])
		if err != nil {
			return nil, fmt.Errorf("invalid lower bound in range %s: %v", r, err)
		}
		upper, err := strconv.Atoi(bounds[1])
		if err != nil {
			return nil, fmt.Errorf("invalid upper bound in range %s: %v", r, err)
		}

		outputRanges[i] = [2]int{lower, upper}
	}

	return outputRanges, nil
}
