package main

type CellsRepository interface {
	SetRandomValues()
	GetCell(x int, y int) (Cell, error)
	SetCell(x int, y int, isAlive bool) error
	GetNeighbors(x int, y int) ([]Cell, error)
	SetNeighbors()
	ApplyRules() error
	CreateGridForDisplay() []byte
	CountCellsThatAreAlive() int
}
