package grid

import (
	"fmt"
	"time"

	"github.com/uhuaha/game-of-life/cells"
)

type grid struct {
	dx      int
	dy      int
	cells   cells.CellsRepository
	display []byte
}

func NewGrid(x int, y int) GridRepository {
	cells := cells.NewCells(x, y)
	cells.SetRandomValues()
	display := cells.CreateGridForDisplay()
	return &grid{dx: x, dy: y, cells: cells, display: display}
}

func (g *grid) GetDimensions() (int, int) {
	return g.dx, g.dy
}

func (g *grid) CalculateNexGeneration() error {
	err := g.cells.ApplyRules()
	if err != nil {
		return err
	}
	g.display = g.cells.CreateGridForDisplay()
	return nil
}

func (g *grid) Draw() {
	fmt.Print("\033[H\033[2J")                 // Clear screen
	fmt.Print("\x0c", string(g.display), "\n") // Print frame
	time.Sleep(1 * time.Second)                // Delay between frames
}
