package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/fabian4/kavicat/event"
)

var (
	App    fyne.App
	Window fyne.Window
)

func Init() {
	App = app.NewWithID("github.com/fabian4/kavicat")
	//App.Settings().SetTheme(customer.NewTheme())
	Window = App.NewWindow("kavicat")
	Window.Resize(fyne.NewSize(1000, 600))
	Window.CenterOnScreen()
	Window.SetMainMenu(NewMenu())

	bottom := container.NewVBox(widget.NewSeparator(), NewBottom())

	split := container.NewHSplit(NewConnection(), NewWork())
	split.Offset = 0.15

	event.Register("inform", Inform)

	content := container.NewBorder(nil, bottom, nil, nil, split)

	Window.SetContent(content)

	Window.ShowAndRun()
}

func GetWindow() fyne.Window {
	return Window
}
