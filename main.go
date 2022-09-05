package main

import (
	"awesomeProject/AutomataDisplay"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
	"time"
)

func run() {
	rand.Seed(time.Now().Unix())
	autoDisp := AutomataDisplay.NewAutomataDisplay(1920/5, 1080/5)
	autoDisp.Run()
}

func main() {
	pixelgl.Run(run)
}
