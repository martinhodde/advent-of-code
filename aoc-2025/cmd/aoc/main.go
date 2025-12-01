package main

import (
	"flag"
	"fmt"
	"log"

	"aoc-2025/internal/registry"
	_ "aoc-2025/internal/solutions" // Ensure solutions are registered
	"aoc-2025/internal/util"
)

func main() {
	day := flag.Int("day", 0, "day number (1-12)")
	part := flag.Int("part", 0, "part number (1 or 2)")
	inputPath := flag.String("input", "", "custom input file path")
	flag.Parse()

	if *day < 1 || *day > 12 {
		log.Fatalf("invalid day: %d", *day)
	}
	if *part != 1 && *part != 2 {
		log.Fatalf("invalid part: %d", *part)
	}

	solver := registry.Lookup(*day, *part)
	if solver == nil {
		log.Fatalf("no solver registered for day %d part %d", *day, *part)
	}

	input := util.LoadInput(*day, *inputPath)
	fmt.Println(solver(input))
}
