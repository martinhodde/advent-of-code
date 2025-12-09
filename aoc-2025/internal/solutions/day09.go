package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
	"math"
)

func init() {
	registry.Register(9, 1, SolveDay9Part1)
	registry.Register(9, 2, SolveDay9Part2)
}

func SolveDay9Part1(input []string) string {
	coordinates := ParseTileCoordinates(input)
	maxRectangleArea := MaxRectangleArea(coordinates)
	return fmt.Sprintf("The largest possible rectangle area using any two coordinates as opposite corners is %d", maxRectangleArea)
}

func SolveDay9Part2(input []string) string {
	return ""
}

// MaxRectangleArea computes the area of the largest rectangle that can be formed
// using any two of the provided coordinates as opposite corners.
func MaxRectangleArea(coordinates [][2]int) int {
	maxArea := 0
	n := len(coordinates)
	for i := range n {
		for j := i + 1; j < n; j++ {
			x1, y1 := coordinates[i][0], coordinates[i][1]
			x2, y2 := coordinates[j][0], coordinates[j][1]
			area := int((math.Abs(float64(x2-x1)) + 1) * (math.Abs(float64(y2-y1)) + 1))
			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}

// ParseTileCoordinates parses a list of strings representing 2D tile
// coordinates in the format "x,y".
func ParseTileCoordinates(input []string) [][2]int {
	var positions [][2]int
	for _, line := range input {
		var x, y int
		if _, err := fmt.Sscanf(line, "%d,%d", &x, &y); err != nil {
			continue // Skip invalid lines (none present in input)
		}
		positions = append(positions, [2]int{x, y})
	}

	return positions
}
