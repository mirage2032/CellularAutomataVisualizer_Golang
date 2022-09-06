package CellAutomata

type Limits struct {
	Limits [][]int
}

func NewLimits(max int, chunks int) Limits {
	var array = make([]int, max-1)
	for i := 1; i < max; i++ {
		array[i-1] = i
	}
	var result [][]int

	for i := 0; i < chunks; i++ {
		min := (i * len(array) / chunks)
		max := ((i + 1) * len(array)) / chunks

		result = append(result, array[min:max])

	}

	var limits [][]int
	for _, j := range result {
		limits = append(limits, []int{j[0], j[len(j)-1] + 1})
	}

	return Limits{limits}
}
