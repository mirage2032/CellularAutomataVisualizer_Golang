package main

import (
	"awesomeProject/AutomataDisplay"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	autoDisp := AutomataDisplay.NewAutomataDisplay(1920*3, 1080*3)
	autoDisp.Run()
}

func main() {
	pixelgl.Run(run)
}
