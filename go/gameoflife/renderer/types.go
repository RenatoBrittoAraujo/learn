package renderer

import "github.com/renatobrittoaraujo/go-gameoflife/gol"

type RenderData struct {
	GameState    gol.GameState
	ScreenWidth  int
	ScreenHeight int
}

type MouseAction struct {
	x       int
	y       int
	w       int
	h       int
	onClick func(renderData *RenderData)
}
