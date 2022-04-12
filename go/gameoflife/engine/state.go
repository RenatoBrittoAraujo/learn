package engine

import "github.com/renatobrittoaraujo/go-gameoflife/gof"

type Engine struct {
	gameState      *gof.GameState
	screenWidth    int
	screenHeight   int
	configs        *Config
	renderDataChan chan *RenderData
}

type RenderData struct {
	gameState    gof.GameState
	screenWidth  int
	screenHeight int
}

type RenderConfigs struct {
	shouldRender   bool `json:"shouldRender"`
	screenWidthPx  int  `json:"screenWidthPx"`
	screenHeightPx int  `json:"screenHeightPx"`
}

type GOFConfigs struct {
	width      int      `json:"width"`
	height     int      `json:"height"`
	startState [][]bool `json:"startState"`
	iterations int      `json:"iterations"`
}

type Config struct {
	GOFData    GOFConfigs    `json:"GOFData"`
	renderData RenderConfigs `json:"renderData"`
}
