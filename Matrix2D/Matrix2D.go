package Matrix2D

type Matrix2D[T any] struct {
	width, height int
	Mat           [][]T
}

func NewMatrix2D[T any](width, height int) *Matrix2D[T] {
	matrix := make([][]T, width)
	for y := range matrix {
		matrix[y] = make([]T, height)
	}
	return &Matrix2D[T]{width, height, matrix}
}

func (mat *Matrix2D[T]) Clear() {
	*mat = *NewMatrix2D[T](mat.width, mat.height)
}
