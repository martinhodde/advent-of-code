package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
)

// Symbol representing a paper roll in the grid
const paperRoll = '@'

// Directions for 8 neighboring cells
var directions = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func init() {
	registry.Register(4, 1, SolveDay4Part1)
	registry.Register(4, 2, SolveDay4Part2)
}

func SolveDay4Part1(input []string) string {
	grid := ParseGrid(input)
	numPaperRolls := NumPaperRollsAccessibleByForklift(grid)
	return fmt.Sprintf("The number of paper rolls accessible by forklift is %d", numPaperRolls)
}

func SolveDay4Part2(input []string) string {
	return ""
}

// NumPaperRollsAccessibleByForklift counts the number of paper rolls in the grid
// that can be accessed by a forklift.
func NumPaperRollsAccessibleByForklift(grid [][]rune) int {
	count := 0
	for y := range len(grid) {
		for x := range len(grid[0]) {
			if grid[y][x] == paperRoll && IsAccessibleByForklift(grid, x, y) {
				count++
			}
		}
	}

	return count
}

// IsAccessibleByForklift checks if a paper roll at position (x, y) can be accessed by a forklift
// based on the surrounding paper rolls. If there are 4 or more adjacent paper rolls, it is not accessible.
func IsAccessibleByForklift(grid [][]rune, x, y int) bool {
	rows, cols := len(grid), len(grid[0])
	if x < 0 || x >= cols || y < 0 || y >= rows {
		return false
	}

	paperRollCount := 0
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < cols && ny >= 0 && ny < rows && grid[ny][nx] == paperRoll {
			paperRollCount++
			if paperRollCount >= 4 {
				return false
			}
		}
	}

	return true
}

// ParseGrid converts the input strings into a 2D grid of runes.
func ParseGrid(input []string) [][]rune {
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}

	return grid
}
