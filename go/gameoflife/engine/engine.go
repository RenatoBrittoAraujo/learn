package engine

import (
	"fmt"

	"github.com/renatobrittoaraujo/go-gameoflife/gof"
	"github.com/renatobrittoaraujo/go-gameoflife/renderer"
)

func CreateEngine(w int, h int) *Engine {
	configs, _ := loadConfigs()

	width := w
	height := h
	tableWidth := w
	tableHeight := h
	var table *gof.Table

	if configs != nil {
		width = configs.renderData.screenWidthPx
		height = configs.renderData.screenHeightPx
		table = (*gof.Table)(&configs.GOFData.startState)
		tableWidth = configs.GOFData.width
		tableHeight = configs.GOFData.height
	}

	renderDataChan := make(chan *RenderData)

	return &Engine{
		gameState:      gof.CreateGameState(tableWidth, tableHeight, table),
		screenWidth:    width,
		screenHeight:   height,
		configs:        configs,
		renderDataChan: renderDataChan,
	}
}

func (e *Engine) RunGame() *gof.GameState {
	fmt.Println("Initializing game...")

	if e.configs.renderData.shouldRender {
		fmt.Println("Renderer created...")
		renderer.InitializeGame()
	}

	fmt.Println("Game has initialized, starting main loop")
	endGameState := e.runLoop()

	fmt.Println("Main loop returned, exiting...")

	// TODO: Logging of results
	// TODO: Abort catch and save
	return endGameState
}

func (e *Engine) runLoop() *gof.GameState {
	if e.configs != nil && e.configs.GOFData.iterations > -1 {
		e.runIterationsLoop(e.configs.GOFData.iterations)
	} else {
		e.runRenderLoop()
	}

	return e.gameState
}

func (e *Engine) runIterationsLoop(iterations int) {
	currentIterations := 0

	for currentIterations < iterations {
		gof.UpdateGameState(e.gameState)
		currentIterations++
	}
}

func (e *Engine) runRenderLoop() {
	currentIterations := 0

	for {
		gof.UpdateGameState(e.gameState)
		currentIterations++
		renderData := e.getRenderData()
		e.renderDataChan <- renderData
	}
}

func (e *Engine) getRenderData() *RenderData {
	return &RenderData{
		gameState:    *e.gameState,
		screenWidth:  e.screenWidth,
		screenHeight: e.screenHeight,
	}
}
