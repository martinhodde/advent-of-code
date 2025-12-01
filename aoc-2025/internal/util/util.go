package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// LoadInput reads the input file and returns a slice of strings (one per line)
func LoadInput(day int, override string) []string {
	path := override
	if path == "" {
		path = fmt.Sprintf("inputs/day%02d.txt", day)
	}
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to load input: %v", err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}
	return lines
}
