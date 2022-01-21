package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewBottom() fyne.CanvasObject {

	bottom := widget.NewLabel("Built by ===")

	return bottom
}
