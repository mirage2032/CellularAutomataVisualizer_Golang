package CellAutomata

import "awesomeProject/Matrix2D"

type cellAutomataBase struct {
	width, height int
	Mat           *Matrix2D.Matrix2D[bool]
}

type CellAutomata interface {
	Step()
	InitMT()
	StepMT()
	Randomize()
	GetMatrix() *Matrix2D.Matrix2D[bool]
	W() int
	H() int
}

func (automata *cellAutomataBase) W() int {
	return automata.width
}
func (automata *cellAutomataBase) H() int {
	return automata.height
}

func (automata *cellAutomataBase) GetMatrix() *Matrix2D.Matrix2D[bool] {
	return automata.Mat
}

func newCellAutomataBase(width, height int) cellAutomataBase {
	automataBase := new(cellAutomataBase)
	automataBase.width = width
	automataBase.height = height
	automataBase.Mat = Matrix2D.NewMatrix2D[bool](width, height)
	return *automataBase
}
