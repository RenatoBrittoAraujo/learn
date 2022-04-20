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
	tileImage       *ebiten.Image
	tableWidth      int
	tableHeight     int
)

func setWidthAndHeight(renderData *RenderData) {
	tableWidth = renderData.ScreenWidth
	tableHeight = renderData.ScreenHeight - bottomGUIHeight
}

func drawTable(renderData *RenderData, screen *ebiten.Image) {
	if tableWidth == 0 || tableHeight == 0 {
		setWidthAndHeight(renderData)
	}

	screen.Fill(backgroundColor)

	drawLines(renderData, screen)
	drawColumns(renderData, screen)
	drawTiles(renderData, screen)
}

func drawLines(renderData *RenderData, screen *ebiten.Image) {
	width := tableWidth - 2*tablePadding

	initializeImageIfNotAlready(&lineImage, width*renderScale, 1)

	op := &ebiten.DrawImageOptions{}
	lineCount := renderData.GameState.H

	translateDelta := (float64(tableHeight) - 2*float64(tablePadding)) / (float64(lineCount))

	op.GeoM.Translate(float64(tablePadding)*renderScale, float64(tablePadding)*renderScale)
	for i := 0; i < lineCount+1; i++ {
		screen.DrawImage(lineImage, op)
		op.GeoM.Translate(0, translateDelta*renderScale)
	}
}

func drawColumns(renderData *RenderData, screen *ebiten.Image) {
	height := tableHeight - 2*tablePadding

	initializeImageIfNotAlready(&columnImage, 1, height*renderScale)

	op := &ebiten.DrawImageOptions{}
	columnCount := renderData.GameState.W

	translateDelta := (float64(tableWidth) - 2*float64(tablePadding)) / (float64(columnCount))

	op.GeoM.Translate(float64(tablePadding)*renderScale, float64(tablePadding)*renderScale)
	for i := 0; i < columnCount+1; i++ {
		screen.DrawImage(columnImage, op)
		op.GeoM.Translate(translateDelta*renderScale, 0)
	}
}

func drawTiles(renderData *RenderData, screen *ebiten.Image) {
	w := renderData.GameState.W
	h := renderData.GameState.H

	tileWidth := (float64(tableWidth)-2*float64(tablePadding))/float64(w) + 0.05
	tileHeight := (float64(tableHeight)-2*float64(tablePadding))/float64(h) + 0.05

	initializeImageIfNotAlready(&tileImage, int(tileWidth*renderScale), int(tileHeight*renderScale))
	for y, row := range *renderData.GameState.Table {
		for x, isPainted := range row {
			if !isPainted {
				continue
			}
			op := &ebiten.DrawImageOptions{}
			xTranslate := float64(tablePadding) + float64(x)*float64(tileWidth)
			yTranslate := float64(tablePadding) + float64(y)*float64(tileHeight)
			op.GeoM.Translate(xTranslate*renderScale, yTranslate*renderScale)
			screen.DrawImage(tileImage, op)
		}
	}
}

func initializeImageIfNotAlready(targetImage **ebiten.Image, w, h int) {
	if (*targetImage) == nil {
		newImage, err := ebiten.NewImage(w, h, ebiten.FilterDefault)
		if err != nil || newImage == nil {
			return
		}
		(*targetImage) = newImage
		(*targetImage).Fill(color.Black)
	}
}
