package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func Inform(args ...string) {

	success := dialog.NewInformation(args[0], args[1], GetWindow())
	success.Resize(fyne.NewSize(200, 150))
	success.Show()
}
