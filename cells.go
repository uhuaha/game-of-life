package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cells struct {
	cell [][]Cell
	dx   int
	dy   int
}

func NewCells(x int, y int) CellsRepository {
	c := make([][]Cell, x)
	for i := range c {
		c[i] = make([]Cell, y)
	}
	return &cells{cell: c, dx: x, dy: y}
}

// SetRandomValues excludes the outer border of cells; the outer border will always be false (i.e. dead)
func (c *cells) SetRandomValues() {
	for i := 1; i < c.dx-1; i++ {
		for j := 1; j < c.dy-1; j++ {
			c.cell[i][j].isAlive = randomBool()
		}
	}
	c.SetNeighbors()
}

func (c *cells) SetCell(x int, y int, isAlive bool) error {
	if x < 0 || x > c.dx || y < 0 || y > c.dy {
		return fmt.Errorf("error: provided x and or y are out of range")
	}
	if x == 0 || x == c.dx || y == 0 || y == c.dy {
		return fmt.Errorf("error: outer most boundary of cells cannot be set; they are always false")
	}
	c.cell[x][y].isAlive = isAlive
	return nil
}

func (c *cells) GetCell(x int, y int) (Cell, error) {
	if x < 0 || x > c.dx || y < 0 || y > c.dy {
		return Cell{}, fmt.Errorf("error: provided x and or y are out of range")
	}
	return c.cell[x][y], nil
}

func (c *cells) GetNeighbors(x int, y int) ([]Cell, error) {
	cell, err := c.GetCell(x, y)
	if err != nil {
		return []Cell{}, err
	}
	return cell.neighbors, nil
}

func randomBool() bool {
	rand.Seed(time.Now().UnixNano())
	rBool := rand.Int()%2 == 0
	return rBool
}

func (c *cells) SetNeighbors() {
	for i := 1; i < c.dx-1; i++ {
		for j := 1; j < c.dy-1; j++ {

			if c.cell[i][j].neighbors != nil {
				c.cell[i][j].neighbors = []Cell{}
			}

			// loop over the eight cells around the current cell
			for ii := i - 1; ii <= i+1; ii++ {
				for jj := j - 1; jj <= j+1; jj++ {
					if ii == i && jj == j {
						continue
					}
					c.cell[i][j].neighbors = append(c.cell[i][j].neighbors, c.cell[ii][jj])
				}
			}

		}
	}
}

func (c *cells) ApplyRules() error {
	for i := 1; i < c.dx-1; i++ {
		for j := 1; j < c.dy-1; j++ {
			aliveNeighbors, err := countAliveNeighbors(c.cell[i][j].neighbors)
			if err != nil {
				return err
			}

			if c.cell[i][j].isAlive && (aliveNeighbors < 2 || aliveNeighbors > 3) {
				c.cell[i][j].isAlive = false
			}

			if !c.cell[i][j].isAlive && aliveNeighbors == 3 {
				c.cell[i][j].isAlive = true
			}
		}
	}
	return nil
}

// CreateGridForDisplay is adapted from: https://play.golang.org/p/93MKfvvP-E
func (c *cells) CreateGridForDisplay() []byte {
	width := 2*c.dx + 2  // +2 for right column + '\n'
	height := 2*c.dy + 1 // +1 for bottom row

	g := make([]byte, width*height)

	ii := 0
	for i := 0; i < height; i += 2 {
		jj := 0
		row0 := i * width
		row1 := (i + 1) * width
		for j := 0; j < width-2; j += 2 {
			g[row0+j], g[row0+j+1] = '+', '-'
			if row1+j+1 <= width*height {
				g[row1+j] = '|'
				if ii < c.dx && jj < c.dy {
					if c.cell[ii][jj].isAlive {
						g[row1+j+1] = '*'
					} else {
						g[row1+j+1] = '.' // \u271D
					}
				}
			}
			jj++
		}
		ii++
		g[row0+width-2], g[row0+width-1] = '+', '\n'
		if row1+width < width*height {
			g[row1+width-2], g[row1+width-1] = '|', '\n'
		}
	}

	return g
}

func (c *cells) CountCellsThatAreAlive() int {
	isAlive := 0
	for i := 1; i < c.dx-1; i++ {
		for j := 1; j < c.dy-1; j++ {
			if c.cell[i][j].isAlive {
				isAlive++
			}
		}
	}
	return isAlive
}

func countAliveNeighbors(ce []Cell) (int, error) {
	if len(ce) == 0 || ce == nil {
		return 0, fmt.Errorf("error: neighbors not defined")
	}
	alive := 0
	for _, c := range ce {
		if c.isAlive {
			alive++
		}
	}
	return alive, nil
}
