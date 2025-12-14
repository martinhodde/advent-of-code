package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
)

// Symbols that appear in the grid
const paperRoll = '@'
const emptySpace = '.'

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
	grid := parseGrid(input)
	numPaperRolls := len(paperRollsAccessibleByForklift(grid))
	return fmt.Sprintf("The number of paper rolls accessible by forklift is %d", numPaperRolls)
}

func SolveDay4Part2(input []string) string {
	grid := parseGrid(input)
	numPaperRolls := numPaperRollsRemoved(grid)
	return fmt.Sprintf("The total number of paper rolls removed by the forklift is %d", numPaperRolls)
}

// numPaperRollsRemoved calculates the total number of paper rolls that can be removed
// from the grid by repeatedly removing accessible paper rolls until none remain.
func numPaperRollsRemoved(grid [][]rune) int {
	removedCount := 0
	paperRolls := paperRollsAccessibleByForklift(grid)
	for len(paperRolls) > 0 {
		removedCount += len(paperRolls)
		removePaperRolls(grid, paperRolls)
		paperRolls = paperRollsAccessibleByForklift(grid)
	}

	return removedCount
}

// removePaperRolls removes the paper rolls at the specified locations
// from the grid by marking them as empty.
func removePaperRolls(grid [][]rune, locations [][2]int) {
	for _, loc := range locations {
		x, y := loc[0], loc[1]
		grid[x][y] = emptySpace
	}
}

// paperRollsAccessibleByForklift computes the coordinates of the paper rolls
// in the grid that can be accessed by a forklift.
func paperRollsAccessibleByForklift(grid [][]rune) [][2]int {
	locations := [][2]int{}
	for x := range len(grid) {
		for y := range len(grid[0]) {
			if grid[x][y] == paperRoll && isAccessibleByForklift(grid, x, y) {
				locations = append(locations, [2]int{x, y})
			}
		}
	}

	return locations
}

// isAccessibleByForklift checks if a paper roll at position (x, y) can be accessed by a forklift
// based on the surrounding paper rolls. If there are 4 or more adjacent paper rolls, it is not accessible.
func isAccessibleByForklift(grid [][]rune, x, y int) bool {
	rows, cols := len(grid), len(grid[0])
	if x < 0 || x >= rows || y < 0 || y >= cols {
		return false
	}

	paperRollCount := 0
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == paperRoll {
			paperRollCount++
			if paperRollCount >= 4 {
				return false
			}
		}
	}

	return true
}

// parseGrid converts the input strings into a 2D grid of runes.
func parseGrid(input []string) [][]rune {
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}

	return grid
}
