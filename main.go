package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	x, y, err := parseArguments()
	if err != nil {
		log.Fatalf("%v", err)
	}
	grid := NewGrid(x, y)
	grid.Draw()

	// Calculate next five generations of the grid
	for i := 0; i < 5; i++ {
		grid.CalculateNexGeneration()
		grid.Draw()
	}
}

func parseArguments() (int, int, error) {
	args := os.Args
	// args = []string{"", "8", "8"}

	x, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Error: Grid dimension x cannot be converted to an integer: %v", err)
	}

	y, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, 0, fmt.Errorf("Error: Grid dimension y cannot be converted to an integer: %v", err)
	}

	if x < 3 || y < 3 {
		return 0, 0, fmt.Errorf("Error: Grid dimension x and/or y is too small. x and y must be at least >= 3.")
	}

	if x > 20 || y > 20 {
		return 0, 0, fmt.Errorf("Error: Grid dimension x and/or y is too big for display. x and y must be <= 20.")
	}

	return x, y, nil
}
