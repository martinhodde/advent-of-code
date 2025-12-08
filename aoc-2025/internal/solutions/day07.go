package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
	"slices"
)

// Manifold characters
const start = 'S'
const splitter = '^'

func init() {
	registry.Register(7, 1, SolveDay7Part1)
	registry.Register(7, 2, SolveDay7Part2)
}

func SolveDay7Part1(input []string) string {
	totalSplits := TotalBeamSplits(input)
	return fmt.Sprintf("The beam is split %d times", totalSplits)
}

func SolveDay7Part2(input []string) string {
	return ""
}

// TotalBeamSplits calculates the total number of beam splits that occur
// as a beam traverses through the manifold represented by the input grid.
func TotalBeamSplits(input []string) int {
	totalSplits := 0
	beamLocs := []int{GetBeamStartLocation(input)}

	// Splitters are located on every other row starting from the third row (index 2)
	for i := 2; i < len(input); i += 2 {
		splitterLocations := GetSplitterLocations(input[i])
		numSplits, newBeamLocs := FindBeamSplitsAndNewBeamLocations(beamLocs, splitterLocations)
		totalSplits += numSplits
		beamLocs = newBeamLocs
	}

	return totalSplits
}

// FindBeamSplitsAndNewBeamLocations identifies the number of beam splits that occur
// at the current row of the manifold and computes the new beam locations after the splits.
func FindBeamSplitsAndNewBeamLocations(beamLocations []int, splitterLocations []int) (int, []int) {
	numSplits := 0
	var newBeamLocations []int

	// Retain beam locations that do not intersect with any splitter locations
	for _, beamLoc := range beamLocations {
		if !slices.Contains(splitterLocations, beamLoc) {
			newBeamLocations = append(newBeamLocations, beamLoc)
		}
	}

	// Check each splitter location to see if it intersects with any current beam locations
	for _, splitLoc := range splitterLocations {
		if slices.Contains(beamLocations, splitLoc) {
			numSplits++
			newBeamLocations = append(newBeamLocations, splitLoc-1)
			newBeamLocations = append(newBeamLocations, splitLoc+1)
		}
	}

	return numSplits, newBeamLocations
}

// GetSplitterLocations locates all column indices marked by the splitter character
// in the given row of the manifold.
func GetSplitterLocations(row string) []int {
	var splitterLocations []int
	for col, char := range row {
		if char == splitter {
			splitterLocations = append(splitterLocations, col)
		}
	}

	return splitterLocations
}

// GetBeamStartLocation locates the starting column index marked by the starting
// character in the first line of the input grid.
func GetBeamStartLocation(input []string) int {
	for col, char := range input[0] {
		if char == start {
			return col
		}
	}

	return -1
}
