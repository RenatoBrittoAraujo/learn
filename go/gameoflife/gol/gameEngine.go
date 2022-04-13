package gol

import (
	"time"
)

func getNeighborsCount(gm *GameState, x, y int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			cx := x + i
			cy := y + j

			if cx < 0 || cx >= gm.W-1 {
				continue
			}
			if cy < 0 || cy >= gm.H-1 {
				continue
			}

			if (*gm.Table)[cx][cy] {
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

func UpdateGameState(gm *GameState) {
	time.Sleep(time.Duration(gm.maxTickSpeedMS * int(time.Millisecond)))

	nextTable := getEmptyTable(gm.W, gm.H)

	for row := 0; row < gm.H-1; row++ {
		(*nextTable)[row] = make([]bool, gm.W)

		for column := 0; column < gm.W-1; column++ {
			isPopulated := (*gm.Table)[row][column]
			neighborsCount := getNeighborsCount(gm, row, column)

			(*nextTable)[row][column] = getActivation(neighborsCount, isPopulated)
		}
	}

	gm.Table = nextTable
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
		300,
	}
}
