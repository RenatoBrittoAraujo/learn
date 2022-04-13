package engine

import (
	"fmt"

	"github.com/renatobrittoaraujo/go-gameoflife/gol"
	"github.com/renatobrittoaraujo/go-gameoflife/renderer"
)

func CreateEngine(w int, h int) *Engine {
	configs, _ := loadConfigs()

	width := w
	height := h
	tableWidth := w
	tableHeight := h
	var table gol.Table

	if configs != nil {
		width = configs.RenderData.ScreenWidthPx
		height = configs.RenderData.ScreenHeightPx
		table = configs.GOLData.StartState
		tableWidth = configs.GOLData.Width
		tableHeight = configs.GOLData.Height
	}

	renderDataChan := make(chan *renderer.RenderData)

	return &Engine{
		gameState:      gol.CreateGameState(tableWidth, tableHeight, table),
		screenWidth:    width,
		screenHeight:   height,
		configs:        configs,
		renderDataChan: renderDataChan,
	}
}

func (e *Engine) RunGame() *gol.GameState {

	go func() {
		fmt.Println("Starting main loop...")
		e.runLoop()
		fmt.Println("Main loop returned, exiting...")
	}()

	if e.configs != nil && e.configs.RenderData.ShouldRender {
		fmt.Println("Creating renderer...")
		initialRenderData := e.getRenderData()
		renderer.InitializeGame(e.renderDataChan, initialRenderData)
	}

	// TODO: Logging of results

	return e.gameState
}

func (e *Engine) runLoop() *gol.GameState {
	if e.configs != nil && e.configs.GOLData.Iterations > -1 {
		e.runIterationsLoop(e.configs.GOLData.Iterations)
	} else {
		e.runRenderLoop()
	}
	close(e.renderDataChan)

	return e.gameState
}

func (e *Engine) runIterationsLoop(iterations int) {
	currentIterations := 0

	for currentIterations < iterations {
		gol.UpdateGameState(e.gameState)
		currentIterations++
	}
}

func (e *Engine) runRenderLoop() {
	currentIterations := 0

	for {
		renderData := e.getRenderData()
		e.renderDataChan <- renderData
		gol.UpdateGameState(e.gameState)
		currentIterations++
	}
}

func (e *Engine) getRenderData() *renderer.RenderData {
	return &renderer.RenderData{
		GameState:    *e.gameState,
		ScreenWidth:  e.screenWidth,
		ScreenHeight: e.screenHeight,
	}
}
