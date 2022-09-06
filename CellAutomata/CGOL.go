package CellAutomata

import (
	"awesomeProject/Matrix2D"
	"math/rand"
	"runtime"
	"sync"
)

type cgol struct {
	cellAutomataBase
	neighbourCount *Matrix2D.Matrix2D[uint8]
	limits         Limits
	threads        int
}

func NewCGOL(width, height int) *cgol {
	automata := new(cgol)
	automata.cellAutomataBase = newCellAutomataBase(width, height)
	automata.neighbourCount = Matrix2D.NewMatrix2D[uint8](width, height)
	automata.threads = runtime.NumCPU()
	automata.limits = NewLimits(height-1, automata.threads)
	return automata
}

func (automata *cgol) Randomize() {
	for x := 1; x < automata.width-1; x++ {
		for y := 1; y < automata.height-1; y++ {
			cell := &automata.Mat.Mat[x][y]
			if rand.Intn(2) == 1 {
				*cell = true
			} else {
				*cell = false
			}
		}
	}
}

func (automata *cgol) countCellNeighbours(x, y int) {
	currentCellNeighbours := &automata.neighbourCount.Mat[x][y]
	*currentCellNeighbours = 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if x == i && y == j { //don't count itself as neighbour
				continue
			}
			if automata.Mat.Mat[i][j] {
				*currentCellNeighbours++
			}
		}
	}
}

func (automata *cgol) countMTNeighbours(y1 int, y2 int) {
	for i := 1; i < automata.width-1; i++ {
		for j := y1; j < y2; j++ {
			automata.countCellNeighbours(i, j)
		}
	}
}

func (automata *cgol) countAllNeighbours() {
	automata.countMTNeighbours(1, automata.height-1)
}

func (automata *cgol) updateCell(x, y int) {
	cell := &automata.Mat.Mat[x][y]
	neighbours := automata.neighbourCount.Mat[x][y]
	if *cell == true && (neighbours < 2 || neighbours > 3) {
		*cell = false
	} else if *cell == false && neighbours == 3 {
		*cell = true
	}
}

func (automata *cgol) updateMTCells(y1 int, y2 int) {
	for i := 1; i < automata.width-1; i++ {
		for j := y1; j < y2; j++ {
			automata.updateCell(i, j)
		}
	}
}

func (automata *cgol) updateCells() {
	automata.updateMTCells(1, automata.height-1)
}

func (automata *cgol) Step() {
	automata.countAllNeighbours()
	automata.updateCells()
}

func (automata *cgol) StepMT() {
	if automata.height-2 < automata.threads {
		automata.Step()
		return
	}
	//Split the Y axis (height) into segments and assign a segment to each thread.
	//Each thread will calculate the neighbours of the cells assigned to it and then update the cells accordingly
	var wg sync.WaitGroup

	for _, limit := range automata.limits.Limits {
		wg.Add(1)
		limit := limit
		go func() {
			defer wg.Done()
			automata.countMTNeighbours(limit[0], limit[1])
		}()
	}
	wg.Wait()

	for _, limit := range automata.limits.Limits {
		wg.Add(1)
		limit := limit
		go func() {
			defer wg.Done()
			go automata.updateMTCells(limit[0], limit[1])
		}()
	}
	wg.Wait()
}
