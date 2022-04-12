package main

import "github.com/renatobrittoaraujo/go-gameoflife/engine"

func main() {
	engine := engine.CreateEngine(500, 500)
	engine.RunGame()
}
