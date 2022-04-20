package gol

type Table [][]bool

type GameState struct {
	W                 int
	H                 int
	Table             *Table
	maxTickSpeedMS    int
	SimulationRunning bool
}
