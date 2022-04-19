package renderer

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	renderScale = 2
)

var (
	initialRenderData *RenderData
	renderDataChan    chan *RenderData
)

type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	renderData, ok := <-renderDataChan

	if ok && screen != nil {
		drawTable(renderData, screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(float64(outsideWidth) * renderScale), int(float64(outsideHeight) * renderScale)
}

func InitializeGame(_renderDataChan chan *RenderData, _initialRenderData *RenderData) {
	renderDataChan = _renderDataChan
	initialRenderData = _initialRenderData
	game := &Game{}

	ebiten.SetWindowSize(initialRenderData.ScreenWidth, initialRenderData.ScreenHeight)
	ebiten.SetWindowTitle("Game of Life")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
