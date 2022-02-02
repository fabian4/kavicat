package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func Success(text string) {
	success := dialog.NewInformation("Success", text, GetWindow())
	success.Resize(fyne.NewSize(400, 300))
	success.Show()
}

func Fail(text string) {
	fail := dialog.NewInformation("Fail", text, GetWindow())
	fail.Resize(fyne.NewSize(400, 300))
	fail.Show()
}
