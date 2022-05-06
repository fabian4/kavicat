package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func ConnectionSuccess(args ...string) {
	remindInfo := dialog.NewInformation(args[0], args[1], GetWindow())
	remindInfo.Resize(fyne.NewSize(200, 150))
	remindInfo.Show()
}

func ConnectionExist(args ...string) {
	remindInfo := dialog.NewInformation("Connection Exist", args[0], GetWindow())
	remindInfo.Resize(fyne.NewSize(200, 150))
	remindInfo.Show()
}

func ConnectionFail(args ...string) {
	remindInfo := dialog.NewInformation(args[0], args[1], GetWindow())
	remindInfo.Resize(fyne.NewSize(200, 150))
	//switch args[2] {
	//case "redis":
	//	dialog.SetOnClosed(newConnectionForRedis)
	//	break
	//}
	remindInfo.Show()
}

func OperationSuccess(args ...string) {
	remindInfo := dialog.NewInformation(args[0], args[1], GetWindow())
	remindInfo.Resize(fyne.NewSize(200, 150))
	remindInfo.Show()
}

func OperationFail(args ...string) {
	remindInfo := dialog.NewInformation("Operation Fail", args[0], GetWindow())
	remindInfo.Resize(fyne.NewSize(200, 150))
	remindInfo.Show()
}
