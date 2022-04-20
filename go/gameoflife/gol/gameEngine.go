package gol

import (
	"time"
)

const (
	tickSpeed = 500
)

func (gm *GameState) getNeighborsCount(x, y int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			cx := x + i
			cy := y + j

			if cx < 0 || cx >= gm.W {
				continue
			}
			if cy < 0 || cy >= gm.H {
				continue
			}

			if (*gm.Table)[cy][cx] {
				count++
			}
		}
	}

	return count
}

func getActivation(neighborsCount int, isPopulated bool) bool {
	// Any live cell with two or three live neighbours survives.
	if isPopulated && (neighborsCount == 2 || neighborsCount == 3) {
		return true
	}

	// Any dead cell with three live neighbours becomes a live cell.
	if !isPopulated && neighborsCount == 3 {
		return true
	}

	// All other live cells die in the next generation. Similarly, all other dead cells stay dead.
	return false
}

func (gm *GameState) UpdateGameState() {
	time.Sleep(time.Duration(gm.maxTickSpeedMS * int(time.Millisecond)))

	if !gm.SimulationRunning {
		return
	}

	nextTable := getEmptyTable(gm.W, gm.H)

	for row := 0; row < gm.H; row++ {
		(*nextTable)[row] = make([]bool, gm.W)

		for column := 0; column < gm.W; column++ {
			isPopulated := (*gm.Table)[row][column]
			neighborsCount := gm.getNeighborsCount(column, row)

			(*nextTable)[row][column] = getActivation(neighborsCount, isPopulated)
		}
	}

	(*gm.Table) = *nextTable
}

func CreateGameState(w, h int, initialTable Table) *GameState {
	var table *Table

	if initialTable == nil {
		table = getEmptyTable(w, h)
	} else {
		table = &initialTable
	}

	return &GameState{
		w,
		h,
		table,
		tickSpeed,
		true,
	}
}

func (gm *GameState) SwitchTile(x, y int) {
	if x < 0 || x >= gm.W || y < 0 || y >= gm.H {
		panic("Invalid input for tile update")
	}
	(*gm.Table)[y][x] = !(*gm.Table)[y][x]
}

func (gm *GameState) SwitchSimulationStatus() {
	gm.SimulationRunning = !gm.SimulationRunning
}
