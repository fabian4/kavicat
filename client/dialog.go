package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func ConnectionSuccess(args ...string) {
	dialog := dialog.NewInformation(args[0], args[1], GetWindow())
	dialog.Resize(fyne.NewSize(200, 150))
	dialog.Show()
}

func ConnectionExist(args ...string) {
	dialog := dialog.NewInformation("Connection Exist", args[0], GetWindow())
	dialog.Resize(fyne.NewSize(200, 150))
	dialog.Show()
}

func ConnectionFail(args ...string) {
	dialog := dialog.NewInformation(args[0], args[1], GetWindow())
	dialog.Resize(fyne.NewSize(200, 150))
	switch args[2] {
	case "redis":
		dialog.SetOnClosed(newConnectionForRedis)
		break
	}
	dialog.Show()
}
