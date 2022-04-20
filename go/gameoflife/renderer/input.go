package renderer

import (
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
			for _, action := range mouseActions {
				executeActionIfClicked(x, y, &action, renderData)
			}

			isClicked = true
		}
	} else {
		isClicked = false
	}
}

func executeActionIfClicked(x, y int, action *MouseAction, renderData *RenderData) bool {
	inXRange := x >= action.x && x <= action.x+action.w
	inYRange := y >= action.y && y <= action.y+action.h

	if inXRange && inYRange {
		action.onClick(renderData)
		return true
	}
	return false
}

func setMouseAction(name string, value MouseAction) {
	value.x *= renderScale
	value.y *= renderScale
	value.w *= renderScale
	value.h *= renderScale
	if _, ok := mouseActions[name]; !ok {
		mouseActions[name] = value
	}
}
