package pt

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/fogleman/gg"
)

type Game struct {
	Name   string
	Width  int
	Height int
	Level  LevelIface
	Image  *image.RGBA
	Debug  bool
}

func New(name string, w, h int, level LevelIface) *Game {
	return &Game{
		Name:   name,
		Width:  w,
		Height: h,
		Level:  level,
		Debug:  true,
		Image:  image.NewRGBA(image.Rect(0, 0, w, h)),
	}
}

func Run(name string, w, h int, level LevelIface) {
	g := New(name, w, h, level)

	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle(name)
	ebiten.SetTPS(60)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	return g.Level.Update()
}

func setBackground(g *Game, drawContext *gg.Context) {
	drawContext.SetColor(color.Black)
	drawContext.DrawRectangle(0, 0, float64(g.Width), float64(g.Height))
	drawContext.Fill()
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawContext := gg.NewContextForRGBA(g.Image)
	setBackground(g, drawContext)
	g.Level.Draw(drawContext)
	screen.WritePixels(g.Image.Pix)
	if g.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Width, g.Height
}
