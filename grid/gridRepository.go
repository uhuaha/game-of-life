package grid

type GridRepository interface {
	GetDimensions() (int, int)
	Draw()
	CalculateNexGeneration() error
}
