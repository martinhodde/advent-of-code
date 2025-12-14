package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
	"strings"
)

// Relevant device names
const youDevice = "you"
const svrDevice = "svr"
const targetDevice = "out"
const dacDevice = "dac"
const fftDevice = "fft"

func init() {
	registry.Register(11, 1, SolveDay11Part1)
	registry.Register(11, 2, SolveDay11Part2)
}

func SolveDay11Part1(input []string) string {
	connections := parseDeviceConnections(input)
	numPaths := numPaths(youDevice, targetDevice, connections)
	return fmt.Sprintf("The number of distinct paths from %s to %s is %d", youDevice, targetDevice, numPaths)
}

func SolveDay11Part2(input []string) string {
	connections := parseDeviceConnections(input)
	numPaths := numPathsWithDACAndFFT(svrDevice, targetDevice, connections)
	return fmt.Sprintf("The number of distinct paths from %s to %s is %d", svrDevice, targetDevice, numPaths)
}

// numPaths computes the number of distinct paths from the start device to the target device
// in a network of devices represented as a DAG.
func numPaths(start, target string, connections map[string][]string) int {
	// Cache results in memoization table
	memo := make(map[string]int)

	var dp func(string) int
	dp = func(node string) int {
		if node == target {
			return 1
		}
		if val, exists := memo[node]; exists {
			return val
		}

		count := 0
		for _, neighbor := range connections[node] {
			count += dp(neighbor)
		}

		memo[node] = count
		return count
	}

	return dp(start)
}

// numPathsWithDACAndFFT computes the number of distinct paths from the start device to the target device
// in a network of devices represented as a DAG. This version assumes that the "dac" and "fft" nodes
// must be included in each valid path, in any order.
func numPathsWithDACAndFFT(start, target string, connections map[string][]string) int {
	// Cache results in memoization table
	type state struct {
		node       string
		dacVisited bool
		fftVisited bool
	}
	memo := make(map[state]int)

	var dfs func(string, bool, bool) int
	dfs = func(current string, dacVisited, fftVisited bool) int {
		if current == target {
			if dacVisited && fftVisited {
				return 1
			}
			return 0
		}

		if current == dacDevice {
			dacVisited = true
		}
		if current == fftDevice {
			fftVisited = true
		}

		key := state{current, dacVisited, fftVisited}
		if val, exists := memo[key]; exists {
			return val
		}

		count := 0
		for _, neighbor := range connections[current] {
			count += dfs(neighbor, dacVisited, fftVisited)
		}

		memo[key] = count
		return count
	}

	return dfs(start, false, false)
}

// parseDeviceConnections parses a list of device connection strings into a map
// where each key is a device and the value is a list of devices it connects to.
// The connections are unidirectional as specified in the input, resulting in a DAG.
func parseDeviceConnections(input []string) map[string][]string {
	connections := make(map[string][]string)
	for _, line := range input {
		parts := strings.Fields(line)
		device := strings.TrimSuffix(parts[0], ":")
		connections[device] = parts[1:]
	}

	return connections
}
