package renderer

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

var (
	backgroundColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
	frameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
	tablePadding    = 10
	lineImage       *ebiten.Image
	columnImage     *ebiten.Image
)

func drawTable(renderData *RenderData, screen *ebiten.Image) {
	// paint screen over
	screen.Fill(backgroundColor)

	drawLines(renderData, screen)

	// draw horizontal lines

	// for each filled up square in table
	//   draw it
}

func drawLines(renderData *RenderData, screen *ebiten.Image) {
	if lineImage == nil {

		width := renderData.ScreenWidth - 2*tablePadding

		newLineImage, err := ebiten.NewImage(1, width, ebiten.FilterDefault)
		if err != nil || newLineImage == nil {
			return
		}
		lineImage = newLineImage
		lineImage.Fill(color.Black)
	}

	op := &ebiten.DrawImageOptions{}
	lineCount := renderData.GameState.H + 1

	translateDelta := (float64(renderData.ScreenHeight) - 2*float64(tablePadding)) / (float64(lineCount))

	op.GeoM.Translate(float64(tablePadding), float64(tablePadding))
	for i := 0; i < lineCount; i++ {
		op.GeoM.Translate(0, translateDelta*float64(i))
		screen.DrawImage(lineImage, op)
	}
}

func drawColumns(w, h int) {

}

func drawTile(w, h, x, y int) {

}
