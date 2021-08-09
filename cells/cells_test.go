package cells

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatCellsCanBeReset(t *testing.T) {
	cells := NewCells(5, 5)
	cells.SetRandomValues()

	cell, err := cells.GetCell(2, 2)
	assert.NoError(t, err)
	isAliveAfterCreation := cell.isAlive

	if isAliveAfterCreation {
		err = cells.SetCell(2, 2, false)
	} else {
		err = cells.SetCell(2, 2, true)
	}
	assert.NoError(t, err)

	cell, err = cells.GetCell(2, 2)
	isAliveAfterReset := cell.isAlive

	assert.NotEqual(t, isAliveAfterCreation, isAliveAfterReset)
}

func TestThatItSetsNeighbors(t *testing.T) {
	cells := NewCells(3, 3)
	cells.SetRandomValues()
	err := cells.SetCell(2, 2, true)
	assert.NoError(t, err)

	neighbors, err := cells.GetNeighbors(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, neighbors)
	assert.Len(t, neighbors, 8)
	assert.Equal(t, 8, countBools(neighbors, false))
	assert.Equal(t, 0, countBools(neighbors, true))
}

func TestThatItSetsNeighborsCorrectly(t *testing.T) {
	cells := NewCells(5, 5)
	cells.SetRandomValues()

	err := cells.SetCell(1, 1, true)
	assert.NoError(t, err)
	err = cells.SetCell(1, 2, false)
	assert.NoError(t, err)
	err = cells.SetCell(1, 3, true)
	assert.NoError(t, err)
	err = cells.SetCell(2, 1, true)
	assert.NoError(t, err)
	err = cells.SetCell(2, 3, false)
	assert.NoError(t, err)
	err = cells.SetCell(3, 1, false)
	assert.NoError(t, err)
	err = cells.SetCell(3, 2, true)
	assert.NoError(t, err)
	err = cells.SetCell(3, 3, true)
	assert.NoError(t, err)

	cells.SetNeighbors()
	neighbors, err := cells.GetNeighbors(2, 2)
	assert.NoError(t, err)
	assert.NotNil(t, neighbors)
	assert.Len(t, neighbors, 8)
	assert.Equal(t, 3, countBools(neighbors, false))
	assert.Equal(t, 5, countBools(neighbors, true))
}

func TestThatDeadCellBecomesAlive(t *testing.T) {
	cells := NewCells(4, 4)
	cells.SetRandomValues()

	err := cells.SetCell(1, 1, false)
	assert.NoError(t, err)
	err = cells.SetCell(1, 2, true)
	assert.NoError(t, err)
	err = cells.SetCell(2, 1, true)
	assert.NoError(t, err)
	err = cells.SetCell(2, 2, true)
	assert.NoError(t, err)

	cells.SetNeighbors()

	err = cells.ApplyRules()
	assert.NoError(t, err)

	centerCell, err := cells.GetCell(1, 1)
	assert.NoError(t, err)

	assert.Equal(t, true, centerCell.isAlive)
}

func TestThatCellDies(t *testing.T) {
	cells := NewCells(4, 4)
	cells.SetRandomValues()

	err := cells.SetCell(1, 1, true)
	assert.NoError(t, err)
	err = cells.SetCell(1, 2, false)
	assert.NoError(t, err)
	err = cells.SetCell(2, 1, false)
	assert.NoError(t, err)
	err = cells.SetCell(2, 2, true)
	assert.NoError(t, err)

	cells.SetNeighbors()

	err = cells.ApplyRules()
	assert.NoError(t, err)

	centerCell, err := cells.GetCell(1, 1)
	assert.NoError(t, err)

	assert.Equal(t, false, centerCell.isAlive)

	// call ApplyRules() a second time
	err = cells.ApplyRules()
	assert.NoError(t, err)

	centerCell, err = cells.GetCell(2, 2)
	assert.NoError(t, err)

	assert.Equal(t, false, centerCell.isAlive)
}

func countBools(cells []Cell, boolean bool) int {
	count := 0
	for _, c := range cells {
		if c.isAlive == boolean {
			count++
		}
	}
	return count
}
