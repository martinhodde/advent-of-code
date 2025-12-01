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
	zeroCount := ZeroCount(input)
	return fmt.Sprintf("Dial landed on position 0 a total of %d times", zeroCount)
}

func SolvePart2(input []string) string {
	return "Day 1 Part 2 not implemented"
}

// ZeroCount takes a list of rotation instructions and returns the number of times
// the dial lands on position 0.
func ZeroCount(rotations []string) int {
	count := 0
	currPos := startPos
	for _, move := range rotations {
		dir, clicks, err := ParseMove(move)
		if err != nil {
			continue // Ignore inavalid rotations
		}

		currPos = RotateDial(currPos, dir, clicks)
		if currPos == 0 {
			count++
		}
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

// RotateDial calculates the new position of the dial after rotating it.
func RotateDial(currPos, clicks, dir int) int {
	return ((currPos + dir*clicks) + dialSize) % dialSize
}
