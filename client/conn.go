package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewConnection() fyne.CanvasObject {

	var data = []string{"a", "striniiiiiiiiiiiiiiiiiiiiiiiig", "list"}
	connection := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
	return connection
}
