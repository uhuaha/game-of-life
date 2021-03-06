package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/uhuaha/game-of-life/grid"
)

func main() {
	x, y, err := parseArguments()
	if err != nil {
		log.Fatalf("error: Cannot parse arguments: %v", err)
	}
	grid := grid.NewGrid(x, y)
	grid.Draw()

	// Calculate next five generations of the grid
	for i := 0; i < 3; i++ {
		err := grid.CalculateNexGeneration()
		if err != nil {
			log.Fatalf("error: Cannot calculate next grid generation: %v", err)
		}
		grid.Draw()
	}
}

func parseArguments() (int, int, error) {
	args := os.Args

	x, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error: Grid dimension x cannot be converted to an integer: %v", err)
	}

	y, err := strconv.Atoi(args[2])
	if err != nil {
		return 0, 0, fmt.Errorf("error: Grid dimension y cannot be converted to an integer: %v", err)
	}

	if x < 3 || y < 3 {
		return 0, 0, fmt.Errorf("error: Grid dimension x and/or y is too small. x and y must be at least >= 3.")
	}

	if x > 20 || y > 20 {
		return 0, 0, fmt.Errorf("error: Grid dimension x and/or y is too big for display. x and y must be <= 20.")
	}

	return x, y, nil
}
