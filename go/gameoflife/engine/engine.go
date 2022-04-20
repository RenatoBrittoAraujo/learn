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
		tableWidth = len(table[0])
		tableHeight = len(table)
	}

	return &Engine{
		gameState:      gol.CreateGameState(tableWidth, tableHeight, table),
		screenWidth:    width,
		screenHeight:   height,
		configs:        configs,
		renderDataChan: make(chan *renderer.RenderData),
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
		e.renderDataChan <- e.getRenderData()
		renderer.InitializeGame(e.renderDataChan)
	}

	// TODO: Logging of results

	return e.gameState
}

func (e *Engine) runLoop() {
	iterations := -1
	if e.configs != nil && e.configs.GOLData.Iterations > -1 {
		iterations = e.configs.GOLData.Iterations
	}

	currentIterations := 0

	for iterations < 0 || currentIterations < iterations {
		e.updateRenderData()

		e.renderData.GameState.UpdateGameState()
		e.renderDataChan <- e.renderData
		currentIterations++
	}
}

func (e *Engine) updateRenderData() {
	select {
	case renderData, ok := <-e.renderDataChan:
		if ok {
			e.renderData = renderData
		}
	default:
	}
}
func (e *Engine) getRenderData() *renderer.RenderData {
	if e.renderData != nil {
		return e.renderData
	}
	return &renderer.RenderData{
		GameState:    *e.gameState,
		ScreenWidth:  e.screenWidth,
		ScreenHeight: e.screenHeight,
	}
}
