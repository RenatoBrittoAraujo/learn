package gof

type Table [][]bool

type GameState struct {
	w     int
	h     int
	table *Table
}
