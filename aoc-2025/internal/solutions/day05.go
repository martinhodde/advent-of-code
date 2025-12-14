package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func init() {
	registry.Register(5, 1, SolveDay5Part1)
	registry.Register(5, 2, SolveDay5Part2)
}

func SolveDay5Part1(input []string) string {
	freshIngredientIDRanges, err := parseFreshIngredientIDRanges(input)
	if err != nil {
		return "invalid fresh ingredient ID range: " + err.Error()
	}
	availableIngredientIDs, err := parseAvailableIngredientIDs(input)
	if err != nil {
		return "invalid available ingredient ID: " + err.Error()
	}
	numFreshIngredients := numFreshIngredients(freshIngredientIDRanges, availableIngredientIDs)
	return fmt.Sprintf("The number of fresh ingredients available is %d", numFreshIngredients)
}

func SolveDay5Part2(input []string) string {
	freshIngredientIDRanges, err := parseFreshIngredientIDRanges(input)
	if err != nil {
		return "invalid fresh ingredient ID range: " + err.Error()
	}
	totalFresh := totalFreshIngredients(freshIngredientIDRanges)
	return fmt.Sprintf("The total number of fresh ingredients across all ranges is %d", totalFresh)
}

// totalFreshIngredients computes the total number of fresh ingredient IDs
// across all specified ranges of fresh ingredient IDs.
func totalFreshIngredients(freshRanges [][2]int) int {
	total := 0
	for _, r := range mergeIngredientIDRanges(freshRanges) {
		total += r[1] - r[0] + 1
	}
	return total
}

// numFreshIngredients counts how many available ingredient IDs
// fall within the specified ranges of fresh ingredient IDs.
func numFreshIngredients(freshRanges [][2]int, availableIDs []int) int {
	// Merge overlapping fresh ID ranges for efficient searching
	mergedRanges := mergeIngredientIDRanges(freshRanges)
	freshCount := 0

	// Binary search each available ID in the merged fresh ID ranges
	for _, id := range availableIDs {
		// Find insertion point based on upper bound each range
		i := sort.Search(len(mergedRanges), func(i int) bool {
			return mergedRanges[i][1] >= id
		})
		// Check if the ID falls within the identified range
		if i < len(mergedRanges) && mergedRanges[i][0] <= id {
			freshCount++
		}
	}

	return freshCount
}

// mergeIngredientIDRanges merges overlapping or contiguous ranges of ingredient IDs.
func mergeIngredientIDRanges(ranges [][2]int) [][2]int {
	// Sort ranges by start ID
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := [][2]int{ranges[0]}
	for _, curr := range ranges[1:] {
		last := merged[len(merged)-1]
		if curr[0] <= last[1]+1 {
			// Ranges overlap or are contiguous, merge them
			if curr[1] > last[1] {
				merged[len(merged)-1][1] = curr[1]
			}
		} else {
			// No overlap, add new range
			merged = append(merged, curr)
		}
	}

	return merged
}

// parseFreshIngredientIDRanges parses the ranges of fresh ingredient IDs
// from the input, which appear before a blank line.
func parseFreshIngredientIDRanges(input []string) ([][2]int, error) {
	var ranges [][2]int
	for _, line := range input {
		if line == "" {
			break // End of ID ranges section of input
		}

		idRange := strings.Split(line, "-")
		start, err := strconv.Atoi(idRange[0])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(idRange[1])
		if err != nil {
			return nil, err
		}

		ranges = append(ranges, [2]int{start, end})
	}

	return ranges, nil
}

// parseAvailableIngredientIDs parses the list of available ingredient IDs
// from the input, which appears after a blank line.
func parseAvailableIngredientIDs(input []string) ([]int, error) {
	var available []int
	inAvailableSection := false
	for _, line := range input {
		if line == "" {
			// Trigger the start of the available IDs section
			inAvailableSection = true
			continue
		}
		if inAvailableSection {
			id, err := strconv.Atoi(line)
			if err != nil {
				return nil, err
			}
			available = append(available, id)
		}
	}

	return available, nil
}
