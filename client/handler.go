package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/fabian4/kavicat/data"
)

func ConnectionSuccess(args ...string) {
	remindInfo := dialog.NewInformation(args[0], args[1], Window)
	remindInfo.Resize(fyne.NewSize(200, 150))
	remindInfo.Show()
	data.RefreshRedisKeyLists()
}

func ConnectionFail(args ...string) {
	remindInfo := dialog.NewInformation(args[0], args[1], Window)
	remindInfo.Resize(fyne.NewSize(200, 150))
	remindInfo.Show()
}

func OperationSuccess(args ...string) {
	remindInfo := dialog.NewInformation(args[0], args[1], Window)
	remindInfo.Resize(fyne.NewSize(200, 150))
	remindInfo.Show()
}

func OperationFail(args ...string) {
	remindInfo := dialog.NewInformation("Operation Fail", args[0], Window)
	remindInfo.Resize(fyne.NewSize(200, 150))
	remindInfo.Show()
}

func SwitchUI(args ...string) {
	switch args[0] {
	case "Home":
		NewHome()
	case "Redis":
		NewRedisWork()
	case "Badger":
		NewBadgerWork()
	case "LevelDB":
		NewLevelDBWork()
	}
}

func Empty(args ...string) {
	remindInfo := dialog.NewInformation(args[0], "It seems to be an empty database", Window)
	remindInfo.Resize(fyne.NewSize(200, 150))
	remindInfo.Show()
}
