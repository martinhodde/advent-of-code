package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
	"strings"
)

const startDevice = "you"
const targetDevice = "out"

func init() {
	registry.Register(11, 1, SolveDay11Part1)
	registry.Register(11, 2, SolveDay11Part2)
}

func SolveDay11Part1(input []string) string {
	connections := ParseDeviceConnections(input)
	numPaths := NumPaths(startDevice, targetDevice, connections)
	return fmt.Sprintf("The number of distinct paths from %s to %s is %d", startDevice, targetDevice, numPaths)
}

func SolveDay11Part2(input []string) string {
	return ""
}

// NumPaths computes the number of distinct paths from the start device to the target device
// in a network of devices represented as a graph. Each device can be visited at most once per path.
func NumPaths(start, target string, connections map[string][]string) int {
	visited := make(map[string]bool)

	var dfs func(string) int
	dfs = func(current string) int {
		if current == target {
			return 1
		}
		if visited[current] {
			return 0
		}

		count := 0
		visited[current] = true
		for _, neighbor := range connections[current] {
			count += dfs(neighbor)
		}
		visited[current] = false // Backtrack

		return count
	}

	return dfs(start)
}

// ParseDeviceConnections parses a list of device connection strings into a map
// where each key is a device and the value is a list of devices it connects to.
// The connections are unidirectional as specified in the input.
func ParseDeviceConnections(input []string) map[string][]string {
	connections := make(map[string][]string)
	for _, line := range input {
		parts := strings.Fields(line)
		device := strings.TrimSuffix(parts[0], ":")
		connections[device] = parts[1:]
	}

	return connections
}
