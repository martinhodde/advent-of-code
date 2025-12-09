package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
)

type Point struct {
	x, y int
}

type Rectangle struct {
	minX, minY, maxX, maxY int
}

func init() {
	registry.Register(9, 1, SolveDay9Part1)
	registry.Register(9, 2, SolveDay9Part2)
}

func SolveDay9Part1(input []string) string {
	redTiles := ParseTileCoordinates(input)
	maxRectangleArea := MaxRectangleArea(redTiles)
	return fmt.Sprintf("The largest rectangle area using any two red tiles as opposite corners is %d", maxRectangleArea)
}

func SolveDay9Part2(input []string) string {
	redTiles := ParseTileCoordinates(input)
	maxRectangleArea := MaxInscribedRectangleArea(redTiles)
	return fmt.Sprintf("The largest inscribed rectangle area using two red tiles as opposite corners is %d", maxRectangleArea)
}

// MaxRectangleArea computes the area of the largest rectangle that can be formed
// using any two of the provided coordinates as opposite corners.
func MaxRectangleArea(coords [][2]int) int {
	maxArea := 0

	// Exhaustively check all pairs of coordinates for max area
	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			x1, y1 := coords[i][0], coords[i][1]
			x2, y2 := coords[j][0], coords[j][1]

			minX, maxX := min(x1, x2), max(x1, x2)
			minY, maxY := min(y1, y2), max(y1, y2)

			area := (maxX - minX + 1) * (maxY - minY + 1)
			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}

// MaxInscribedRectangleArea finds the largest rectangle inscribed within the polygon formed
// by connecting red tiles (with green tiles), where two opposite corners must be red tiles.
func MaxInscribedRectangleArea(redTiles [][2]int) int {
	maxArea := 0

	// Check all pairs of red tiles as potential rectangle corners
	for i := range redTiles {
		for j := i + 1; j < len(redTiles); j++ {
			x1, y1 := redTiles[i][0], redTiles[i][1]
			x2, y2 := redTiles[j][0], redTiles[j][1]

			minX, maxX := min(x1, x2), max(x1, x2)
			minY, maxY := min(y1, y2), max(y1, y2)

			rect := Rectangle{minX, minY, maxX, maxY}
			potentialArea := (maxX - minX + 1) * (maxY - minY + 1)
			if potentialArea <= maxArea {
				continue // No need to check smaller areas
			}

			if IsRectangleInsidePolygon(rect, redTiles) {
				maxArea = potentialArea
			}
		}
	}

	return maxArea
}

// IsRectangleInsidePolygon checks if an axis-aligned rectangle is fully inside the polygon formed
// by connecting the red tiles (with green tiles).
func IsRectangleInsidePolygon(rect Rectangle, polygon [][2]int) bool {
	// Check that no polygon edge crosses through the rectangle interior
	for i := range polygon {
		j := (i + 1) % len(polygon) // Next vertex, wrapping around
		start := Point{polygon[i][0], polygon[i][1]}
		end := Point{polygon[j][0], polygon[j][1]}

		if IsSegmentIntersectingRectangle(start, end, rect) {
			return false
		}
	}

	return true
}

// IsSegmentIntersectingRectangle checks if an axis-aligned line segment crosses
// through the interior of a rectangle.
func IsSegmentIntersectingRectangle(start, end Point, rect Rectangle) bool {
	// Terminate early if both segment endpoints are outside rectangle bounds
	if (start.x < rect.minX && end.x < rect.minX) || (start.x > rect.maxX && end.x > rect.maxX) ||
		(start.y < rect.minY && end.y < rect.minY) || (start.y > rect.maxY && end.y > rect.maxY) {
		return false
	}

	// Check axis-aligned segments
	if start.x == end.x {
		// Vertical segment is problematic if strictly inside rectangle's x-range
		if start.x > rect.minX && start.x < rect.maxX {
			minSeg, maxSeg := min(start.y, end.y), max(start.y, end.y)
			if minSeg < rect.maxY && maxSeg > rect.minY {
				return true
			}
		}
	} else if start.y == end.y {
		// Horizontal segment is problematic if strictly inside rectangle's y-range
		if start.y > rect.minY && start.y < rect.maxY {
			minSeg, maxSeg := min(start.x, end.x), max(start.x, end.x)
			if minSeg < rect.maxX && maxSeg > rect.minX {
				return true
			}
		}
	}

	return false
}

// ParseTileCoordinates parses input lines in "x,y" format into coordinate pairs.
func ParseTileCoordinates(input []string) [][2]int {
	var positions [][2]int
	for _, line := range input {
		var x, y int
		if _, err := fmt.Sscanf(line, "%d,%d", &x, &y); err != nil {
			continue
		}
		positions = append(positions, [2]int{x, y})
	}
	return positions
}
