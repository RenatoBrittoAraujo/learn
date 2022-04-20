package renderer

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

var (
	isClicked    bool                   = false
	mouseActions map[string]MouseAction = make(map[string]MouseAction)
)

const (
	PAUSE_UNPAUSE = "pause/unpause"
)

func checkInput(renderData *RenderData) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !isClicked {
			x, y := ebiten.CursorPosition()
			for key, action := range mouseActions {
				executed := checkAction(x, y, &action, renderData)
				if executed {
					fmt.Println("Click found for ", key)
				}
			}

			isClicked = true
		}
	} else {
		isClicked = false
	}
}

func checkAction(x, y int, action *MouseAction, renderData *RenderData) bool {
	inXRange := x >= action.x && x <= action.x+action.w
	inYRange := y >= action.y && y <= action.y+action.h

	fmt.Println("CLICK POS = ", x, y)
	fmt.Println("RECT POS = ", action)

	if inXRange && inYRange {
		action.onClick(renderData)
		return true
	}
	return false
}

func setMouseAction(name string, value MouseAction) {
	if _, ok := mouseActions[name]; !ok {
		mouseActions[name] = value
	}
}
