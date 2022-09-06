package main

import (
	"awesomeProject/AutomataDisplay"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
	"time"
)

func run() {
	rand.Seed(time.Now().Unix())
	autoDisp := AutomataDisplay.NewAutomataDisplay(1920, 1080)
	for i := 0; i < 1000; i++ {
		autoDisp.Render()
		autoDisp.HandleInput()
		autoDisp.Automata.StepMT()
	}
}

func main() {
	pixelgl.Run(run)
}
