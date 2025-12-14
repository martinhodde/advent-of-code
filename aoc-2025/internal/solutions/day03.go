package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
	"math"
)

func init() {
	registry.Register(3, 1, SolveDay3Part1)
	registry.Register(3, 2, SolveDay3Part2)
}

func SolveDay3Part1(input []string) string {
	numBatteries := 2
	batteryBanks := parseBatteryBanks(input)
	outputJoltage := totalJoltage(batteryBanks, numBatteries)
	return fmt.Sprintf("The total maximum joltage using %d batteries per bank is %d", numBatteries, outputJoltage)
}

func SolveDay3Part2(input []string) string {
	numBatteries := 12
	batteryBanks := parseBatteryBanks(input)
	outputJoltage := totalJoltage(batteryBanks, numBatteries)
	return fmt.Sprintf("The total maximum joltage using %d batteries per bank is %d", numBatteries, outputJoltage)
}

// totalJoltage calculates the total maximum joltage that can be achieved from
// a collection of battery banks by selecting a specified number of batteries from each bank.
func totalJoltage(banks [][]int, numBatteries int) int {
	totalJoltage := 0
	for _, bank := range banks {
		totalJoltage += maxJoltage(bank, numBatteries)
	}

	return totalJoltage
}

// maxJoltage calculates the maximum joltage that can be achieved by selecting
// a specified number of batteries from the given bank of battery joltages.
func maxJoltage(batteries []int, numBatteriesLeft int) int {
	// Memoization map to cache results
	memo := make(map[[2]int]int)

	var computeMaxJoltage func(int, int) int
	computeMaxJoltage = func(batteryIdx int, numBatteriesLeft int) int {
		if numBatteriesLeft == 0 {
			// Base case: no more batteries to select
			return 0
		}
		if batteryIdx == len(batteries) {
			// Base case: still more batteries to select, but no more
			// batteries available in the bank (invalid state)
			return -math.MaxInt
		}

		// Memoization key: [current battery index, number of batteries left to select]
		key := [2]int{batteryIdx, numBatteriesLeft}
		if val, exists := memo[key]; exists {
			return val
		}

		// Recursively compute the maximum joltage by either skipping or taking the current battery
		skipCurrent := computeMaxJoltage(batteryIdx+1, numBatteriesLeft)
		takeCurrent := batteries[batteryIdx]*int(math.Pow10(numBatteriesLeft-1)) +
			computeMaxJoltage(batteryIdx+1, numBatteriesLeft-1)

		memo[key] = max(skipCurrent, takeCurrent)
		return memo[key]
	}

	return computeMaxJoltage(0, numBatteriesLeft)
}

// parseBatteryBanks converts a slice of strings representing battery banks
// into a slice of slices of integers representing the joltages of the batteries.
func parseBatteryBanks(banks []string) [][]int {
	var batteryBanks [][]int
	for _, bank := range banks {
		var batteryJoltages []int
		for _, b := range bank {
			// Convert rune to int by subtracting '0'
			batteryJoltages = append(batteryJoltages, int(b-'0'))
		}
		batteryBanks = append(batteryBanks, batteryJoltages)
	}

	return batteryBanks
}
