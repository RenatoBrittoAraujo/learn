package renderer

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	renderScale = 2
)

type Game struct {
	currentRenderData *RenderData
	renderDataChan    chan *RenderData
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	g.tryFetchNewRenderData()

	checkInput(g.currentRenderData)

	g.updateRenderData()
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	drawTable(g.currentRenderData, screen)
	drawGUI(g.currentRenderData, screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(float64(outsideWidth) * renderScale), int(float64(outsideHeight) * renderScale)
}

func InitializeGame(renderDataChan chan *RenderData) {
	game := &Game{}
	game.renderDataChan = renderDataChan
	game.currentRenderData = <-renderDataChan

	ebiten.SetWindowSize(game.currentRenderData.ScreenWidth, game.currentRenderData.ScreenHeight)
	ebiten.SetWindowTitle("Game of Life")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) tryFetchNewRenderData() {
	select {
	case newRenderData, ok := <-g.renderDataChan:
		if ok {
			g.currentRenderData = newRenderData
		}
	default:
	}
}

func (g *Game) updateRenderData() {

	select {
	case _, ok := <-g.renderDataChan:
		if !ok {
			g.renderDataChan <- g.currentRenderData
		}
	default:
	}
}
