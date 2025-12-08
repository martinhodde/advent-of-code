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
	totalTimelines := TotalTimelines(input)
	return fmt.Sprintf("The original beam undergoes %d timelines", totalTimelines)
}

// TotalTimelines calculates the total number of distinct beam timelines that result
// from the beam traversing through the manifold represented by the input grid.
func TotalTimelines(input []string) int {
	// Map from beam location to count of timelines at that location
	beamLocs := make(map[int]int)
	beamLocs[GetBeamStartLocation(input)] = 1

	// Splitters are located on every other row starting from the third row (index 2)
	for i := 2; i < len(input); i += 2 {
		splitterLocations := GetSplitterLocations(input[i])
		_, beamLocs = FindBeamSplitsAndNewBeamLocations(beamLocs, splitterLocations)
	}

	// Sum up all timelines across all final beam locations
	totalTimelines := 0
	for _, count := range beamLocs {
		totalTimelines += count
	}

	return totalTimelines
}

// TotalBeamSplits calculates the total number of beam splits that occur
// as the beam traverses through the manifold represented by the input grid.
func TotalBeamSplits(input []string) int {
	totalSplits := 0

	// Establish initial beam location
	beamLocs := make(map[int]int)
	beamLocs[GetBeamStartLocation(input)] = 1

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
func FindBeamSplitsAndNewBeamLocations(beamLocations map[int]int, splitterLocations []int) (int, map[int]int) {
	numSplits := 0
	newBeamLocations := make(map[int]int)

	// Retain beam locations that do not intersect with any splitter locations
	for beamLoc := range beamLocations {
		if !slices.Contains(splitterLocations, beamLoc) {
			newBeamLocations[beamLoc] = beamLocations[beamLoc]
		}
	}

	// Check each splitter location to see if it intersects with any current beam locations
	for _, splitLoc := range splitterLocations {
		if count, exists := beamLocations[splitLoc]; exists {
			// Split each timeline's beam into two new beams at adjacent locations
			newBeamLocations[splitLoc-1] += count
			newBeamLocations[splitLoc+1] += count
			numSplits++
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
