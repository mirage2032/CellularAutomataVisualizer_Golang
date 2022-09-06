package tests

import (
	"awesomeProject/CellAutomata"
	"testing"
)

func BenchmarkStepSinglethreaded(b *testing.B) {
	automata := CellAutomata.NewCGOL(1920, 1080)
	automata.Randomize()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		automata.Step()
	}
}

func BenchmarkStepMultithreaded(b *testing.B) {
	automata := CellAutomata.NewCGOL(1920, 1080)
	automata.Randomize()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		automata.StepMT()
	}
}
