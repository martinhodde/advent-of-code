package solutions

import (
	"aoc-2025/internal/registry"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const epsilon = 1e-6

type Machine struct {
	TargetLightState    uint32
	Buttons             []uint32
	JoltageRequirements []int
}

func init() {
	registry.Register(10, 1, solveDay10Part1)
	registry.Register(10, 2, solveDay10Part2)
}

func solveDay10Part1(input []string) string {
	machines := parseMachineInfo(input)
	totalPresses := buttonPressSum(machines, fewestButtonPressesToTargetLightStates)
	return fmt.Sprintf(
		"The number of button presses required to achieve the target indicator light states for all machines is %d",
		totalPresses,
	)
}

func solveDay10Part2(input []string) string {
	machines := parseMachineInfo(input)
	total := buttonPressSum(machines, fewestButtonPressesToJoltageRequirements)
	return fmt.Sprintf(
		"The number of button presses required to achieve the joltage requirements for all machines is %d",
		total,
	)
}

// buttonPressSum computes the total number of button presses required
// to achieve the target machine states for all machines.
func buttonPressSum(machines []Machine, pressFunc func(Machine) int) int {
	totalPresses := 0
	for _, machine := range machines {
		presses := pressFunc(machine)
		totalPresses += presses
	}

	return totalPresses
}

// fewestButtonPressesToTargetLightStates computes the minimum number of button presses required
// to achieve the target indicator light state for the given machine configuration.
func fewestButtonPressesToTargetLightStates(machine Machine) int {
	type State struct {
		lightState uint32
		numPresses int
	}

	visited := make(map[uint32]bool)
	queue := []State{{0, 0}} // Light state starts at 0 (all off)
	visited[0] = true

	// Perform BFS to find the shortest path to the target state
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.lightState == machine.TargetLightState {
			return current.numPresses
		}

		for _, button := range machine.Buttons {
			newState := current.lightState ^ button
			if !visited[newState] {
				visited[newState] = true
				queue = append(queue, State{newState, current.numPresses + 1})
			}
		}
	}

	return -1 // Target state is unreachable
}

// fewestButtonPressesToJoltageRequirements computes the minimum number of button presses required
// to achieve the joltage requirements specified for the given machine configuration.
func fewestButtonPressesToJoltageRequirements(machine Machine) int {
	// Problem Formulation:
	// We want to minimize Σx_i subject to A·x = b, x ≥ 0, x ∈ ℤ
	// where:
	//   - x is a vector of button press counts (one entry per button)
	//   - A is a binary matrix where A[i][j] = 1 if button j affects joltage i, else 0
	//   - b is the target joltage requirements vector
	//
	// This is an integer linear programming problem which we solve by first using Gaussian elimination to
	// identify free variables, then searching nearby integer assignments to find the optimal solution.
	numJoltages := len(machine.JoltageRequirements)
	numButtons := len(machine.Buttons)

	// Build binary coefficient matrix A
	A := make([][]float64, numJoltages)
	for i := range numJoltages {
		A[i] = make([]float64, numButtons)
		for j := range numButtons {
			if (machine.Buttons[j] & (1 << i)) != 0 {
				A[i][j] = 1.0
			}
		}
	}

	// Build target joltages vector b
	b := make([]float64, numJoltages)
	for i, req := range machine.JoltageRequirements {
		b[i] = float64(req)
	}

	// Solve the system of linear equations with integer constraints
	solution := solveIntegerLinearSystem(A, b, numButtons, numJoltages)

	total := 0
	for _, v := range solution {
		total += v
	}

	return total
}

// solveIntegerLinearSystem solves A·x = b for non-negative integer x with the minimum vector sum.
func solveIntegerLinearSystem(A [][]float64, b []float64, numVars, numConstraints int) []int {
	// Reduce to RREF and identify pivot columns
	aug, pivotCols := gaussianElimination(A, b, numVars, numConstraints)

	// Determine which variables are free vs dependent
	freeVars := identifyFreeVariables(pivotCols, numVars)

	// Search over all integer assignments to free variables
	return searchFreeVariables(aug, pivotCols, freeVars, numVars)
}

// gaussianElimination performs Gaussian elimination with partial pivoting
// to reduce the augmented matrix [A|b] to reduced row echelon form (RREF)
// Returns the resulting matrix and the list of pivot column indices.
func gaussianElimination(A [][]float64, b []float64, numVars, numConstraints int) ([][]float64, []int) {
	// Create augmented matrix [A|b]
	aug := make([][]float64, numConstraints)
	for i := range numConstraints {
		aug[i] = make([]float64, numVars+1)
		copy(aug[i], A[i])
		aug[i][numVars] = b[i]
	}

	pivotCols := make([]int, 0, numConstraints)
	row := 0

	for col := 0; col < numVars && row < numConstraints; col++ {
		// Find pivot (row with largest absolute value in this column)
		pivotRow := row
		maxVal := math.Abs(aug[row][col])
		for i := row + 1; i < numConstraints; i++ {
			if absVal := math.Abs(aug[i][col]); absVal > maxVal {
				maxVal = absVal
				pivotRow = i
			}
		}

		if maxVal < epsilon {
			continue // Skip if column has no pivot
		}

		// Swap rows to bring pivot to current row
		aug[row], aug[pivotRow] = aug[pivotRow], aug[row]

		// Scale pivot row such that pivot = 1
		pivot := aug[row][col]
		for j := range numVars + 1 {
			aug[row][j] /= pivot
		}

		// Eliminate this column in all other rows (reduced row echelon form)
		for i := range numConstraints {
			if i != row {
				factor := aug[i][col]
				for j := range numVars + 1 {
					aug[i][j] -= factor * aug[row][j]
				}
			}
		}

		pivotCols = append(pivotCols, col)
		row++
	}

	return aug, pivotCols
}

// identifyFreeVariables determines which variables are free (non-pivot).
func identifyFreeVariables(pivotCols []int, numVars int) []int {
	isPivot := make(map[int]bool, len(pivotCols))
	for _, col := range pivotCols {
		isPivot[col] = true
	}

	freeVars := make([]int, 0, numVars-len(pivotCols))
	for i := range numVars {
		if !isPivot[i] {
			freeVars = append(freeVars, i)
		}
	}

	return freeVars
}

// searchFreeVariables performs a bounded search over integer assignments
// to free variables, computing dependent variables for each assignment
// and tracking the best valid solution in terms of its minimum vector sum.
func searchFreeVariables(aug [][]float64, pivotCols []int, freeVars []int, numVars int) []int {
	const maxSearchValue = 500 // Max value to try for each free variable

	bestSolution := []int(nil)
	bestSum := math.MaxInt

	var search func(freeIdx int, assignment []int, partialSum int)
	search = func(freeIdx int, assignment []int, partialSum int) {
		if partialSum >= bestSum {
			return
		}

		// All free variables assigned, compute solution
		if freeIdx == len(freeVars) {
			solution := computeSolutionFromFreeVars(aug, pivotCols, freeVars, assignment, numVars)

			if valid, sum := validateSolution(solution); valid && sum < bestSum {
				bestSum = sum
				bestSolution = make([]int, numVars)
				for i, v := range solution {
					bestSolution[i] = int(math.Round(v))
				}
			}
			return
		}

		// Scan through possible values for current free variable
		for val := 0; val <= maxSearchValue; val++ {
			assignment[freeIdx] = val
			search(freeIdx+1, assignment, partialSum+val)
		}
	}

	assignment := make([]int, len(freeVars))
	search(0, assignment, 0)

	return bestSolution
}

// computeSolutionFromFreeVars computes the full solution vector given an assignment to the
// free variables. Uses the RREF to algebraically determine the pivot variables.
func computeSolutionFromFreeVars(aug [][]float64, pivotCols []int, freeVars []int, assignment []int, numVars int) []float64 {
	solution := make([]float64, numVars)
	for i, val := range assignment {
		solution[freeVars[i]] = float64(val)
	}

	// For each pivot row i: x_pivot[i] = aug[i][last] - sum(aug[i][j] * x[j]) for non-pivot j
	for i, pivotCol := range pivotCols {
		if i >= len(aug) {
			break
		}
		val := aug[i][numVars]
		for j := range numVars {
			if j != pivotCol {
				val -= aug[i][j] * solution[j]
			}
		}
		solution[pivotCol] = val
	}

	return solution
}

// validateSolution checks if a solution is valid (all non-negative integers)
// and returns the validity flag and the sum of all variables.
func validateSolution(solution []float64) (bool, int) {
	sum := 0

	for _, v := range solution {
		// Check non-negative
		if v < -epsilon {
			return false, 0
		}
		// Check integer (within tolerance)
		if math.Abs(v-math.Round(v)) > epsilon {
			return false, 0
		}
		sum += int(math.Round(v))
	}

	return true, sum
}

// parseMachineInfo parses the input lines describing the desired state for each machine.
// Each line of input is formatted like the following example:
// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
// where the first section (brackets) is the target on/off states of each indicator light,
// the middle sections (parentheses) are buttons that toggle specific lights/joltages, and
// the last section (curly braces) represents the joltage requirements of each machine.
func parseMachineInfo(input []string) []Machine {
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
			machine.TargetLightState = bitmap
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
