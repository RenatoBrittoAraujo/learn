package main

import (
	"fmt"

	tb "github.com/nsf/termbox-go"
)

type gameData struct {
	started bool
	table   [3][3]byte
	turns   int
}

func drawBoard(table [3][3]byte) {
	fmt.Println("-------------")
	for i := 0; i < 3; i++ {
		fmt.Printf("| ")
		for j := 0; j < 3; j++ {
			fmt.Printf("%c", table[i][j])

			if j < 2 {
				fmt.Printf(" | ")
			} else {
				fmt.Printf(" |\n")
			}
		}
		fmt.Println("-------------")
	}
}

func getIntegerInput() int {
	var input int
	fmt.Scanf("%d", &input)
	return input
}

func promtKeyPress(message string) string {
	fmt.Println(message)
	var input string
	fmt.Scanf("%s", input)
	return input
}

func clearGameData(gm *gameData) {
	gm.table = [3][3]byte{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}
	gm.turns = 0
	gm.started = false
}

func checkTie(table [3][3]byte) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if table[i][j] == ' ' {
				return false
			}
		}
	}

	return true
}

func checkVictory(table [3][3]byte) (bool, byte) {
	// horz
	for i := 0; i < 3; i++ {
		if table[i][0] != ' ' && table[i][0] == table[i][1] && table[i][1] == table[i][2] {
			return true, table[i][0]
		}
	}

	// vert
	for i := 0; i < 3; i++ {
		if table[0][i] != ' ' && table[0][i] == table[1][i] && table[1][i] == table[2][i] {
			return true, table[0][i]
		}
	}

	// diag
	if table[1][1] != ' ' && table[1][1] == table[0][0] && table[1][1] == table[2][2] {
		return true, table[0][0]
	}

	// counter-diag
	if table[1][1] != ' ' && table[0][2] == table[1][1] && table[1][1] == table[2][0] {
		return true, table[1][1]
	}

	return false, ' '
}

func handleGameTurn(gm *gameData) {
	drawBoard(gm.table)
	turn := 'X'
	if gm.turns%2 == 1 {
		turn = 'Y'
	}

	fmt.Println("")
	fmt.Println("It is " + string([]rune{turn}) + "'s turn!")

	var row, column int
	for {
		row = 0
		fmt.Println("Select a number for the ROW (1 - 3)")
		for row > 3 || row < 1 {
			row = getIntegerInput()
		}

		column = 0
		fmt.Println("Select a number for the COLUMN (1 - 3)")
		for column > 3 || column < 1 {
			column = getIntegerInput()
		}

		if gm.table[row-1][column-1] == ' ' {
			break
		} else {
			fmt.Println("=====> ERROR: The selected cell is already in use!", row, column)
			fmt.Println("")
		}
	}
	fmt.Println("")

	gm.table[row-1][column-1] = byte(turn)

	isTie := checkTie(gm.table)
	isVictory, victor := checkVictory(gm.table)
	if isTie || isVictory {
		drawBoard(gm.table)
		fmt.Println("")

		if isTie {
			fmt.Println("Game has tied!")
		}
		if isVictory {
			fmt.Println("Player", string([]byte{victor}), "has won the game!")
		}

		clearGameData(gm)
		return
	}

	gm.turns++
}

func gameLoop(gm *gameData) {
	for {
		tb.Clear(tb.ColorBlack, tb.ColorBlack)
		tb.Flush()
		if !gm.started {
			promtKeyPress("Press any key to start the game")
			gm.started = true
		} else {
			handleGameTurn(gm)
		}
	}

}

func main() {
	fmt.Println("Welcome to Tic Tac Toe!")

	table := [3][3]byte{
		{' ', ' ', ' '},
		{' ', ' ', ' '},
		{' ', ' ', ' '},
	}

	gameData := gameData{
		table:   table,
		started: false,
		turns:   0,
	}

	gameLoop(&gameData)
}
