package engine

import "github.com/renatobrittoaraujo/go-gameoflife/gol"

type Engine struct {
	gameState      *gol.GameState
	screenWidth    int
	screenHeight   int
	configs        *Config
	renderDataChan chan *RenderData
}

type RenderData struct {
	gameState    gol.GameState
	screenWidth  int
	screenHeight int
}

type RenderConfigs struct {
	shouldRender   bool `json:"shouldRender"`
	screenWidthPx  int  `json:"screenWidthPx"`
	screenHeightPx int  `json:"screenHeightPx"`
}

type GOLConfigs struct {
	width      int      `json:"width"`
	height     int      `json:"height"`
	startState [][]bool `json:"startState"`
	iterations int      `json:"iterations"`
}

type Config struct {
	GOLData    GOLConfigs    `json:"GOlData"`
	renderData RenderConfigs `json:"renderData"`
}
