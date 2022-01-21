package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func NewWork() fyne.CanvasObject {

	work := canvas.NewText("content", theme.ForegroundColor())

	return work
}
