package gof

func getNeighborsCount(table *Table, x, y int) int {
	count := 0

	for i := 1; i >= -1; i -= 2 {
		for j := 1; j >= -1; j -= 2 {
			if (*table)[x+j][y+i] {
				count++
			}
		}
	}

	return count
}

func getActivation(neighborsCount int, initialCellState bool) bool {
	// Any live cell with two or three live neighbours survives.
	if initialCellState && (neighborsCount == 2 || neighborsCount == 3) {
		return true
	}

	// Any dead cell with three live neighbours becomes a live cell.
	if !initialCellState && neighborsCount == 3 {
		return true
	}

	// All other live cells die in the next generation. Similarly, all other dead cells stay dead.
	return false
}

func UpdateGameState(gm *GameState) {
	nextTable := getEmptyTable(gm.w, gm.h)

	for row := 0; row < gm.h; row++ {
		(*nextTable)[row] = make([]bool, gm.w)

		for column := 0; column < gm.w; column++ {
			neighborsCount := getNeighborsCount(gm.table, row, column)
			(*nextTable)[row][column] = getActivation(neighborsCount, (*gm.table)[row][column])
		}
	}

	gm.table = nextTable
}

func CreateGameState(w, h int, initialTable *Table) *GameState {
	var table *Table

	if initialTable == nil {
		table = getEmptyTable(w, h)
	} else {
		table = initialTable
	}

	return &GameState{
		w,
		h,
		table,
	}
}
