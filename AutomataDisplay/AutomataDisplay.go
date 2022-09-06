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
		Title:       "Automata",
		Bounds:      pixel.R(0, 0, float64(width), float64(height)),
		VSync:       false,
		Undecorated: true,
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

func (adsp *AutomataDisplay) automataToCanvas() *pixelgl.Canvas {
	//create empty Canvas(black by default)
	pic := pixelgl.NewCanvas(adsp.cfg.Bounds)
	pixels := pic.Pixels()
	//iterate through the automata matrix and if a cell is active, change the pixel's color
	matrix := adsp.automata.GetMatrix().Mat
	for i := 0; i < adsp.automata.W(); i++ {
		for j := 0; j < adsp.automata.H(); j++ {
			if matrix[i][j] == true {
				pixels[(i+j*int(adsp.cfg.Bounds.Max.X))*4] = 255   //red channel
				pixels[(i+j*int(adsp.cfg.Bounds.Max.X))*4+1] = 255 //red channel
				pixels[(i+j*int(adsp.cfg.Bounds.Max.X))*4+2] = 255 //red channel
				pixels[(i+j*int(adsp.cfg.Bounds.Max.X))*4+3] = 255 //alpha channel
			}
		}
	}
	pic.SetPixels(pixels)
	return pic
}

func (adsp *AutomataDisplay) handleInput() {
	if adsp.win.Pressed(pixelgl.KeyEscape) {
		adsp.win.SetClosed(true)
	}
}

func (adsp *AutomataDisplay) Run() {
	for !adsp.win.Closed() {
		adsp.win.Clear(colornames.Black)
		sprite := adsp.automataToCanvas()
		sprite.Draw(adsp.win, pixel.IM.Moved(adsp.win.Bounds().Center()))
		adsp.win.Update()
		adsp.handleInput()
		adsp.automata.StepMT()
	}
}
