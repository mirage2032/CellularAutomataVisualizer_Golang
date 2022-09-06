package main

import (
	"awesomeProject/AutomataDisplay"
	"github.com/faiface/pixel/pixelgl"
	"github.com/fstanis/screenresolution"
	"math/rand"
	"time"
)

func run() {
	rand.Seed(time.Now().Unix())
	resolution := screenresolution.GetPrimary()
	autoDisp := AutomataDisplay.NewAutomataDisplay(resolution.Width, resolution.Height)
	autoDisp.Run()
}

func main() {
	pixelgl.Run(run)
}
