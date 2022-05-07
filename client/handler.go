package client

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/fabian4/kavicat/data"
)

func ConnectionSuccess(args ...string) {
	remindInfo := dialog.NewInformation(args[0], args[1], GetWindow())
	remindInfo.Resize(fyne.NewSize(200, 150))
	remindInfo.Show()
	data.RefreshKeyLists()
}

func ConnectionExist(args ...string) {
	remindInfo := dialog.NewInformation("Connection Exist", args[0], GetWindow())
	remindInfo.Resize(fyne.NewSize(200, 150))
	remindInfo.Show()
}

func ConnectionFail(args ...string) {
	remindInfo := dialog.NewInformation(args[0], args[1], GetWindow())
	remindInfo.Resize(fyne.NewSize(200, 150))
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

func SwitchUI(args ...string) {
	fmt.Println(args[0])
	switch args[0] {
	case "Home":
		NewHome()
	case "Redis":
		NewRedisWork()
	}
}
