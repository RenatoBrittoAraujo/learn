package engine

import (
	"github.com/renatobrittoaraujo/go-gameoflife/gol"
	"github.com/renatobrittoaraujo/go-gameoflife/renderer"
)

type Engine struct {
	gameState      *gol.GameState
	screenWidth    int
	screenHeight   int
	configs        *Config
	renderDataChan chan *renderer.RenderData
	renderData     *renderer.RenderData
}

type RenderConfigs struct {
	ShouldRender   bool `json:"shouldRender"`
	ScreenWidthPx  int  `json:"screenWidthPx"`
	ScreenHeightPx int  `json:"screenHeightPx"`
}

type GOLConfigs struct {
	StartState [][]bool `json:"startState"`
	Iterations int      `json:"iterations"`
}

type Config struct {
	GOLData    GOLConfigs    `json:"golData"`
	RenderData RenderConfigs `json:"renderData"`
}
