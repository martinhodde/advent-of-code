package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// FlipDirection represents the type of flip applied to a gift.
type FlipDirection int

const (
	FlipNone FlipDirection = iota
	FlipHorizontal
	FlipVertical
	FlipBoth
)

// Rotation represents the angle of clockwise rotation
type Rotation int

const (
	Rotate0 Rotation = iota
	Rotate90
	Rotate180
	Rotate270
)

// Tree represents an under-tree region with its dimensions and requested gift counts,
// where GiftCounts[i] is the number of gifts of shape i requested.
type Tree struct {
	Width      int
	Height     int
	GiftCounts []int
}

func init() {
	registry.Register(12, 1, SolveDay12Part1)
}

func SolveDay12Part1(input []string) string {
	gifts := parseGifts(input)
	trees := parseTrees(input)
	numTrees := numAccommodatingTrees(trees, gifts)
	return fmt.Sprintf("The number of trees that can fit their requested gifts is %d", numTrees)
}

// numAccommodatingTrees counts how many under-tree regions can accommodate their requested gifts.
func numAccommodatingTrees(trees []Tree, gifts [][][2]int) int {
	total := 0
	for _, tree := range trees {
		if canFitGiftsUnderTree(tree, gifts) {
			total++
		}
	}

	return total
}

// canFitGiftsUnderTree determines if the tree can accommodate all requested gifts
// in any orientation without overlap.
func canFitGiftsUnderTree(tree Tree, gifts [][][2]int) bool {
	// Precompute all orientations for each gift
	giftOrientations := make([][][][2]int, len(gifts))
	for i, gift := range gifts {
		giftOrientations[i] = generateAllOrientations(gift)
	}

	// Check if we even have enough total space under the tree for all gifts
	totalCellsNeeded := 0
	for i, count := range tree.GiftCounts {
		totalCellsNeeded += count * len(gifts[i])
	}
	if totalCellsNeeded > tree.Width*tree.Height {
		return false
	}

	// Create a grid to represent occupied spaces under the tree
	occupied := make([][]bool, tree.Height)
	for i := range occupied {
		occupied[i] = make([]bool, tree.Width)
	}

	memo := make(map[string]bool)
	return tryPlaceAllGifts(occupied, giftOrientations, tree.GiftCounts, memo)
}

// tryPlaceAllGifts attempts to place all gifts trying different orientations during placement.
func tryPlaceAllGifts(occupied [][]bool, allOrientations [][][][2]int, counts []int, memo map[string]bool) bool {
	if !slices.ContainsFunc(counts, func(x int) bool { return x != 0 }) {
		return true // All gifts successfully placed
	}

	// Check memoization table
	key := gridStateKey(occupied, counts)
	if result, exists := memo[key]; exists {
		return result
	}

	// Try to place the next gift from any gift that still has remaining quantity
	for giftIdx, count := range counts {
		if count > 0 {
			// Try every orientation of this gift
			for _, orientation := range allOrientations[giftIdx] {
				// Try every possible position under the tree
				for startRow := range occupied {
					for startCol := range occupied[0] {
						if canPlaceGiftAt(occupied, orientation, startRow, startCol) {
							// Place the gift
							placeGiftAt(occupied, orientation, startRow, startCol, true)
							counts[giftIdx]--

							// Recursively try to place remaining gifts
							if tryPlaceAllGifts(occupied, allOrientations, counts, memo) {
								memo[key] = true
								return true
							}

							// Backtrack
							counts[giftIdx]++
							placeGiftAt(occupied, orientation, startRow, startCol, false)
						}
					}
				}
			}

			// If we have tried all orientations and positions for this gift and none worked,
			// there is no point in continuing with the other gifts
			memo[key] = false
			return false
		}
	}

	memo[key] = false
	return false
}

// gridStateKey creates a unique key for memoization
func gridStateKey(occupied [][]bool, counts []int) string {
	var sb strings.Builder

	// Occupied grid
	for _, row := range occupied {
		for _, cell := range row {
			if cell {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
	}
	sb.WriteByte('|')

	// Gift counts
	for i, count := range counts {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.Itoa(count))
	}

	return sb.String()
}

// canPlaceGiftAt checks if a gift can be placed at a specific position.
func canPlaceGiftAt(occupied [][]bool, gift [][2]int, startRow, startCol int) bool {
	for _, coord := range gift {
		row := startRow + coord[0]
		col := startCol + coord[1]
		if row < 0 || row >= len(occupied) || col < 0 || col >= len(occupied[0]) {
			return false // Out of bounds
		}
		if occupied[row][col] {
			return false // Space already occupied
		}
	}

	return true
}

// placeGiftAt places or removes a gift at a specific position.
func placeGiftAt(occupied [][]bool, gift [][2]int, startRow, startCol int, place bool) {
	for _, coord := range gift {
		row := startRow + coord[0]
		col := startCol + coord[1]
		occupied[row][col] = place
	}
}

// generateAllOrientations generates all unique orientations of a given gift.
func generateAllOrientations(gift [][2]int) [][][2]int {
	var orientations [][][2]int
	flipDirections := []FlipDirection{FlipNone, FlipHorizontal, FlipVertical, FlipBoth}
	rotations := []Rotation{Rotate0, Rotate90, Rotate180, Rotate270}

	seen := make(map[string]bool)
	for _, flip := range flipDirections {
		flippedGift := flipGift(gift, flip)
		for _, rotation := range rotations {
			rotatedGift := rotateGift(flippedGift, rotation)
			normalizedGift := normalizeGift(rotatedGift)

			// Create a unique key for each gift orientation to avoid duplicates
			var keyParts []string
			for _, coord := range normalizedGift {
				keyParts = append(keyParts, fmt.Sprintf("%d,%d", coord[0], coord[1]))
			}
			key := strings.Join(keyParts, ";")
			if !seen[key] {
				seen[key] = true
				orientations = append(orientations, normalizedGift)
			}
		}
	}

	return orientations
}

// normalizeGift shifts all gift coordinates so the minimum row and column are both 0.
func normalizeGift(gift [][2]int) [][2]int {
	minRow, minCol := gift[0][0], gift[0][1]
	for _, coord := range gift {
		minRow = min(minRow, coord[0])
		minCol = min(minCol, coord[1])
	}

	normalized := make([][2]int, len(gift))
	for i, coord := range gift {
		normalized[i] = [2]int{coord[0] - minRow, coord[1] - minCol}
	}

	return normalized
}

// flipGift flips the gift coordinates based on the specified flip direction.
func flipGift(gift [][2]int, direction FlipDirection) [][2]int {
	flipped := make([][2]int, len(gift))
	for i, coord := range gift {
		row, col := coord[0], coord[1]
		switch direction {
		case FlipNone:
			flipped[i] = [2]int{row, col}
		case FlipHorizontal:
			flipped[i] = [2]int{row, -col}
		case FlipVertical:
			flipped[i] = [2]int{-row, col}
		case FlipBoth:
			flipped[i] = [2]int{-row, -col}
		}
	}

	return flipped
}

// rotateGift rotates the gift coordinates based on the specified rotation angle.
func rotateGift(gift [][2]int, rotation Rotation) [][2]int {
	rotated := make([][2]int, len(gift))
	for i, coord := range gift {
		row, col := coord[0], coord[1]
		switch rotation {
		case Rotate0:
			rotated[i] = [2]int{row, col}
		case Rotate90:
			rotated[i] = [2]int{col, -row}
		case Rotate180:
			rotated[i] = [2]int{-row, -col}
		case Rotate270:
			rotated[i] = [2]int{-col, row}
		}
	}

	return rotated
}

// parseGifts parses the input lines to extract gift coordinates from their respective
// graphical representations within a grid.
func parseGifts(input []string) [][][2]int {
	var gifts [][][2]int
	var currentGift [][2]int
	row := 0

	for _, line := range input {
		if line == "" {
			// End of current gift
			gifts = append(gifts, currentGift)
			row = 0
			continue
		}
		if line[len(line)-1] == ':' {
			// New gift header
			currentGift = [][2]int{}
			row = 0
			continue
		}
		if strings.Contains(line, "x") {
			break // Start of tree inputs, stop parsing gifts
		}

		for col, char := range line {
			if char == '#' {
				currentGift = append(currentGift, [2]int{row, col})
			}
		}
		row++
	}

	return gifts
}

// parseTrees parses the input lines to extract under-tree region dimensions
// and their corresponding requested gift counts.
func parseTrees(input []string) []Tree {
	var trees []Tree

	for _, line := range input {
		if strings.Contains(line, "x") && strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			dimParts := strings.Split(parts[0], "x")
			width, _ := strconv.Atoi(dimParts[0])
			height, _ := strconv.Atoi(dimParts[1])

			var counts []int
			countParts := strings.Fields(parts[1])
			for _, countStr := range countParts {
				count, _ := strconv.Atoi(countStr)
				counts = append(counts, count)
			}

			trees = append(trees, Tree{
				Width:      width,
				Height:     height,
				GiftCounts: counts,
			})
		}
	}

	return trees
}
