package renderer

import "github.com/renatobrittoaraujo/go-gameoflife/gol"

type RenderData struct {
	GameState    gol.GameState
	ScreenWidth  int
	ScreenHeight int
}
