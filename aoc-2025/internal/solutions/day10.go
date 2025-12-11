package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Each machine starts with all lights off
const initLightState uint32 = 0

type Machine struct {
	TargetState         uint32
	Buttons             []uint32
	JoltageRequirements []int
}

func init() {
	registry.Register(10, 1, SolveDay10Part1)
	registry.Register(10, 2, SolveDay10Part2)
}

func SolveDay10Part1(input []string) string {
	machines := ParseMachineInfo(input)
	totalPresses := ButtonPressSum(machines)
	return fmt.Sprintf(
		"The total number of button presses required to achieve the target states for all machines is %d",
		totalPresses,
	)
}

func SolveDay10Part2(input []string) string {
	return ""
}

// ButtonPressSum computes the total number of button presses required
// to achieve the target indicator light states for all machines.
func ButtonPressSum(machines []Machine) int {
	totalPresses := 0
	for _, machine := range machines {
		presses := FewestButtonPressesToTarget(machine)
		totalPresses += presses
	}

	return totalPresses
}

// FewestButtonPressesToTarget computes the minimum number of button presses required
// to achieve the target indicator light state for the given machine configuration.
func FewestButtonPressesToTarget(machine Machine) int {
	type State struct {
		lightState uint32
		presses    int
	}

	visited := make(map[uint32]bool)
	queue := []State{{initLightState, 0}}
	visited[initLightState] = true

	// Perform BFS to find the shortest path to the target state
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.lightState == machine.TargetState {
			return current.presses
		}

		for _, button := range machine.Buttons {
			newState := current.lightState ^ button
			if !visited[newState] {
				visited[newState] = true
				queue = append(queue, State{newState, current.presses + 1})
			}
		}
	}

	return -1 // Target state is unreachable
}

// ParseMachineInfo parses the input lines describing the desired indicator light state
// for each machine. Each line of input is formatted like the following example:
// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
// where the first section (brackets) is the target on/off states of each indicator light,
// the middle sections (parentheses) are buttons that toggle specific lights, and
// the last section (curly braces) represents the joltage requirements of each light.
func ParseMachineInfo(input []string) []Machine {
	var machines []Machine

	// Regex patterns
	targetStatePattern := regexp.MustCompile(`\[([.#]+)\]`)
	buttonPattern := regexp.MustCompile(`\(([0-9,]+)\)`)
	joltagePattern := regexp.MustCompile(`\{([0-9,]+)\}`)

	for _, line := range input {
		machine := Machine{}

		// Extract target light states as bitmap
		if match := targetStatePattern.FindStringSubmatch(line); len(match) > 1 {
			// Convert string to bitmap: '#' -> 1, '.' -> 0
			stateStr := match[1]
			var bitmap uint32
			for i, ch := range stateStr {
				if ch == '#' {
					bitmap |= 1 << i
				}
			}
			machine.TargetState = bitmap
		}

		// Extract buttons as bitmasks
		for _, match := range buttonPattern.FindAllStringSubmatch(line, -1) {
			if len(match) > 1 {
				var buttonMask uint32
				for _, numStr := range strings.Split(match[1], ",") {
					if num, err := strconv.Atoi(strings.TrimSpace(numStr)); err == nil {
						buttonMask |= 1 << num
					}
				}
				machine.Buttons = append(machine.Buttons, buttonMask)
			}
		}

		// Extract joltage requirements
		if match := joltagePattern.FindStringSubmatch(line); len(match) > 1 {
			for _, numStr := range strings.Split(match[1], ",") {
				if num, err := strconv.Atoi(strings.TrimSpace(numStr)); err == nil {
					machine.JoltageRequirements = append(machine.JoltageRequirements, num)
				}
			}
		}

		machines = append(machines, machine)
	}

	return machines
}
