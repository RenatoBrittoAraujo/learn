package gol

type Table [][]bool

type GameState struct {
	w     int
	h     int
	table *Table
}
