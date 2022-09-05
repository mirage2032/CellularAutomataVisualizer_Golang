package AutomataDisplay

import (
	"awesomeProject/CellAutomata"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
)

type AutomataDisplay struct {
	cfg      pixelgl.WindowConfig
	win      *pixelgl.Window
	automata CellAutomata.CellAutomata
}

func NewAutomataDisplay(width, height int) *AutomataDisplay {
	adsp := new(AutomataDisplay)
	adsp.cfg = pixelgl.WindowConfig{
		Title:  "AUTOMATA",
		Bounds: pixel.R(0, 0, float64(width), float64(height)),
	}
	win, err := pixelgl.NewWindow(adsp.cfg)
	if err != nil {
		panic(err)
	}
	adsp.automata = CellAutomata.NewCGOL(width, height)
	adsp.automata.Randomize()
	adsp.win = win
	return adsp
}

func (adsp *AutomataDisplay) automataToSprite() *pixel.Sprite {
	pic := pixel.MakePictureData(adsp.cfg.Bounds)
	matrix := adsp.automata.GetMatrix().Mat
	for i := 0; i < adsp.automata.W(); i++ {
		for j := 0; j < adsp.automata.H(); j++ {
			if matrix[i][j] == true {
				pic.Pix[i+j*adsp.automata.W()] = color.RGBA{R: 255, G: 255, B: 255, A: 255}
			} else {
				pic.Pix[i+j*adsp.automata.W()] = color.RGBA{A: 255}
			}
		}
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	return sprite
}

func (adsp *AutomataDisplay) Run() {
	for !adsp.win.Closed() {
		adsp.win.Clear(colornames.Black)
		sprite := adsp.automataToSprite()
		sprite.Draw(adsp.win, pixel.IM.Moved(adsp.win.Bounds().Center()))
		adsp.win.Update()
		adsp.automata.Step()
	}
}

func (adsp *AutomataDisplay) RunMT() {
	adsp.automata.InitMT()
	for !adsp.win.Closed() {
		adsp.win.Clear(colornames.Black)
		sprite := adsp.automataToSprite()
		sprite.Draw(adsp.win, pixel.IM.Moved(adsp.win.Bounds().Center()))
		adsp.win.Update()
		adsp.automata.Step()
	}
}
