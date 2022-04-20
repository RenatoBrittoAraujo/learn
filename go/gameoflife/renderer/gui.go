package renderer

import (
	"log"

	img "image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	bottomGUIHeight  = 50
	imgCenterPadding = 25
)

var (
	pauseUnpauseButtonImage *ebiten.Image
	pauseImg                *ebiten.Image
	unpauseImg              *ebiten.Image
)

func init() {
	var err error
	pauseUnpauseButtonImage, _, err = ebitenutil.NewImageFromFile("assets/pauseunpause.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	bounds := pauseUnpauseButtonImage.Bounds()

	width := float64(bounds.Max.X - bounds.Min.X)
	height := float64(bounds.Max.Y - bounds.Min.Y)

	pauseImg = pauseUnpauseButtonImage.SubImage(img.Rectangle{
		Min: img.Point{0, 0},
		Max: img.Point{int(width/2) - imgCenterPadding, int(height)},
	}).(*ebiten.Image)

	unpauseImg = pauseUnpauseButtonImage.SubImage(img.Rectangle{
		Min: img.Point{int(width/2) + imgCenterPadding, 0},
		Max: img.Point{int(width), int(height)},
	}).(*ebiten.Image)
}

func drawGUI(renderData *RenderData, screen *ebiten.Image) {
	if renderData.GameState.SimulationRunning {
		drawPauseButton(renderData, screen)
	} else {
		drawUnpauseButton(renderData, screen)
	}

	drawUpdateSpeedSlider(renderData, screen)
	drawColumnsCounter(renderData, screen)
	drawRowsCounter(renderData, screen)
}

func drawPauseButton(renderData *RenderData, screen *ebiten.Image) {
	drawBottomImage(renderData, screen, pauseImg)
}

func drawUnpauseButton(renderData *RenderData, screen *ebiten.Image) {
	drawBottomImage(renderData, screen, unpauseImg)
}

func drawBottomImage(renderData *RenderData, screen *ebiten.Image, image *ebiten.Image) {
	bounds := image.Bounds()

	width := float64(bounds.Max.X - bounds.Min.X)
	height := float64(bounds.Max.Y - bounds.Min.Y)

	op := &ebiten.DrawImageOptions{}

	percentualChange := height / (bottomGUIHeight - float64(tablePadding))

	if percentualChange > 1 {
		op.GeoM.Scale(renderScale/percentualChange, renderScale/percentualChange)
		width /= percentualChange
		height /= percentualChange
	}

	xpos := float64(tablePadding)
	ypos := float64(renderData.ScreenHeight) - bottomGUIHeight

	op.GeoM.Translate(xpos*renderScale, ypos*renderScale)

	screen.DrawImage(image, op)

	setMouseAction(
		PAUSE_UNPAUSE,
		MouseAction{
			x: int(xpos),
			y: int(ypos),
			w: int(width),
			h: int(height),
			onClick: func(renderData *RenderData) {
				renderData.GameState.SwitchSimulationStatus()
			},
		},
	)
}

func drawUpdateSpeedSlider(renderData *RenderData, screen *ebiten.Image) {

}

func drawColumnsCounter(renderData *RenderData, screen *ebiten.Image) {

}

func drawRowsCounter(renderData *RenderData, screen *ebiten.Image) {

}
