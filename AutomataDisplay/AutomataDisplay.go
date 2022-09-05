package AutomataDisplay

import (
	"awesomeProject/CellAutomata"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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
		VSync:  false,
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
	//create empty PictureData(black by default)
	pic := pixel.MakePictureData(adsp.cfg.Bounds)
	//iterate through the automata matrix and if a cell is active, change the pixel's color
	matrix := adsp.automata.GetMatrix().Mat
	for i := 0; i < adsp.automata.W(); i++ {
		for j := 0; j < adsp.automata.H(); j++ {
			if matrix[i][j] == true {
				pic.Pix[(i + j*adsp.automata.W())] = colornames.White
			}
		}
	}
	//create sprite and return it
	sprite := pixel.NewSprite(pic, pic.Bounds())
	return sprite
}

func (adsp *AutomataDisplay) Run() {
	for !adsp.win.Closed() {
		adsp.win.Clear(colornames.Black)
		sprite := adsp.automataToSprite()
		sprite.Draw(adsp.win, pixel.IM.Moved(adsp.win.Bounds().Center()))
		adsp.win.Update()
		adsp.automata.StepMT()
	}
}
