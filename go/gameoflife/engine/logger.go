package engine

import "github.com/renatobrittoaraujo/go-gameoflife/gol"

type Logs struct {
	logs []Log
}

type Log struct {
	timestamp  string
	iterations int
	gameState  gol.GameState
}

func LogGameState(gm *gol.GameState) {
	// if !file, create
	// else open file

	// append log

	// close file
}
