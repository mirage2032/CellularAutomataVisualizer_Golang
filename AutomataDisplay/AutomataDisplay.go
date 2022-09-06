package AutomataDisplay

import (
	"awesomeProject/CellAutomata"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type AutomataDisplay struct {
	cfg      pixelgl.WindowConfig
	Win      *pixelgl.Window
	Automata CellAutomata.CellAutomata
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
	adsp.Automata = CellAutomata.NewCGOL(width, height)
	adsp.Automata.Randomize()
	adsp.Win = win
	return adsp
}

func (adsp *AutomataDisplay) automataToCanvas() *pixelgl.Canvas {
	//create empty Canvas(black by default)
	pic := pixelgl.NewCanvas(adsp.cfg.Bounds)
	pixels := pic.Pixels()
	//iterate through the Automata matrix and if a cell is active, change the pixel's color
	matrix := adsp.Automata.GetMatrix().Mat
	for i := 0; i < int(adsp.cfg.Bounds.Max.X); i++ {
		for j := 0; j < int(adsp.cfg.Bounds.Max.Y); j++ {
			if i == 0 || j == 0 || i == int(adsp.cfg.Bounds.Max.X)-1 || j == int(adsp.cfg.Bounds.Max.Y)-1 {
				pixels[(i+j*int(adsp.cfg.Bounds.Max.X))*4] = 255   //red channel
				pixels[(i+j*int(adsp.cfg.Bounds.Max.X))*4+1] = 0   //red channel
				pixels[(i+j*int(adsp.cfg.Bounds.Max.X))*4+2] = 0   //red channel
				pixels[(i+j*int(adsp.cfg.Bounds.Max.X))*4+3] = 255 //alpha channel
			}
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

func (adsp *AutomataDisplay) HandleInput() {
	if adsp.Win.Pressed(pixelgl.KeyEscape) {
		adsp.Win.SetClosed(true)
	}
}

func (adsp *AutomataDisplay) Render() {
	adsp.Win.Clear(colornames.Black)
	sprite := adsp.automataToCanvas()
	sprite.Draw(adsp.Win, pixel.IM.Moved(adsp.Win.Bounds().Center()))
	adsp.Win.Update()
}

func (adsp *AutomataDisplay) Run() {
	for !adsp.Win.Closed() {
		adsp.Render()
		adsp.HandleInput()
		adsp.Automata.StepMT()
	}
}
