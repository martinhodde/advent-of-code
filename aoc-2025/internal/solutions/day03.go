package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
)

func init() {
	registry.Register(3, 1, SolveDay3Part1)
	registry.Register(3, 2, SolveDay3Part2)
}

func SolveDay3Part1(input []string) string {
	outputJoltage := TotalJoltage(ParseBatteryBanks(input))
	return fmt.Sprintf("The total maximum joltage is %d", outputJoltage)
}

func SolveDay3Part2(input []string) string {
	return ""
}

// TotalJoltage calculates the total maximum joltage
// that can be achieved from a collection of battery banks.
func TotalJoltage(banks [][]int) int {
	totalJoltage := 0
	for _, bank := range banks {
		totalJoltage += MaxJoltage(bank)
	}

	return totalJoltage
}

// MaxJoltage calculates the maximum joltage that can be achieved
// by selecting two batteries from the given slice of battery joltages.
func MaxJoltage(batteries []int) int {
	maxJoltage := 0
	for left := 0; left < len(batteries)-1; left++ {
		for right := left + 1; right < len(batteries); right++ {
			joltage := 10*batteries[left] + batteries[right]
			maxJoltage = max(joltage, maxJoltage)
		}
	}

	return maxJoltage
}

// ParseBatteryBanks converts a slice of strings representing battery banks
// into a slice of slices of integers representing the joltages of the batteries.
func ParseBatteryBanks(banks []string) [][]int {
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
