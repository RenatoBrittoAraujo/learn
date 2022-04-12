package engine

import "github.com/renatobrittoaraujo/go-gameoflife/gof"

type Logs struct {
	logs []Log
}

type Log struct {
	timestamp  string
	iterations int
	gameState  gof.GameState
}

func LogGameState(gm *gof.GameState) {
	// if !file, create
	// else open file

	// append log

	// close file
}
