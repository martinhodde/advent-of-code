package solutions

import (
	"aoc-2025/internal/registry"
	"errors"
	"fmt"
	"strconv"
)

// Dial constants
const dialSize = 100
const startPos = 50

func init() {
	registry.Register(1, 1, SolvePart1)
	registry.Register(1, 2, SolvePart2)
}

func SolvePart1(input []string) string {
	onlyCountDirect := true
	zeroCount := ZeroCount(input, onlyCountDirect)
	return fmt.Sprintf("Dial landed directly on position 0 a total of %d times", zeroCount)
}

func SolvePart2(input []string) string {
	onlyCountDirect := false
	zeroCount := ZeroCount(input, onlyCountDirect)
	return fmt.Sprintf("Dial encountered position 0 a total of %d times", zeroCount)
}

// ZeroCount takes a list of rotation instructions and returns the number of times
// the dial encounters position 0. If onlyCountDirect is true, it counts only direct landings on 0.
func ZeroCount(rotations []string, onlyCountDirect bool) int {
	count := 0
	currPos := startPos

	for _, move := range rotations {
		dir, clicks, err := ParseMove(move)
		if err != nil {
			continue // Just ignore inavalid rotations
		}

		newPos := RotateDial(currPos, dir, clicks)

		if onlyCountDirect {
			if newPos == 0 {
				count++
			}
		} else {
			// Calculate the minimum number of clicks it would take to reach position 0 going in the given direction
			// This is equivalent to the smallest k ≥ 1 such that currPos + dir*k ≡ 0 (mod dialSize)
			minClicks := ((-dir*currPos)%dialSize + dialSize) % dialSize
			if minClicks == 0 {
				// If already at position 0, the next encounter would be after a full rotation
				minClicks = dialSize
			}

			if clicks >= minClicks {
				// Add 1 for the first encounter, then count additional full rotations
				count += 1 + (clicks-minClicks)/dialSize
			}
		}

		currPos = newPos
	}

	return count
}

// ParseMove takes a move instruction string (e.g., "L10" or "R5") and returns the direction
// and number of clicks as integers.
func ParseMove(move string) (int, int, error) {
	if len(move) < 2 {
		return 0, 0, errors.New("move instruction too short")
	}

	var dir int
	switch move[0] {
	case 'L':
		dir = -1
	case 'R':
		dir = 1
	default:
		return 0, 0, errors.New("invalid move instruction: must start with 'L' or 'R'")
	}

	clicks, err := strconv.Atoi(move[1:])
	if err != nil || clicks < 0 {
		return 0, 0, errors.New("invalid number of steps in move instruction")
	}

	return dir, clicks, nil
}

// RotateDial calculates the new position of the dial after rotation by a given number of clicks
// in a specified direction, starting from the current position.
func RotateDial(currPos, dir, clicks int) int {
	return (currPos + dir*clicks + dialSize) % dialSize
}
