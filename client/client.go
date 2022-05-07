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

	RegisterEvents()

	App = app.NewWithID("github.com/fabian4/kavicat")
	//App.Settings().SetTheme(customer.NewTheme())
	Window = App.NewWindow("kavicat")
	Window.Resize(fyne.NewSize(1000, 600))
	Window.CenterOnScreen()
	//Window.SetMainMenu(NewMenu())

	//bottom := container.NewVBox(widget.NewSeparator(), NewBottom())
	//
	//split := container.NewHSplit(NewList(), NewWork())
	//split.Offset = 0.2
	//
	//content := container.NewBorder(nil, bottom, nil, nil, split)

	NewHome()

	Window.ShowAndRun()
}

func GetWindow() fyne.Window {
	return Window
}

func NewHome() {
	buttons := container.NewVBox(
		widget.NewButton("Establish connection for Redis", newConnectionForRedis),
		widget.NewButton("Establish connection for Badger", func() {

		}),
		widget.NewButton("Establish connection for LevelDB", func() {

		}),
	)
	content := container.NewCenter(buttons)
	Window.SetContent(content)
}

func RegisterEvents() {
	event.Register("switchUI", SwitchUI)
	event.Register("connection_fail", ConnectionFail)
	event.Register("connection_exist", ConnectionExist)
	event.Register("connection_success", ConnectionSuccess)
	event.Register("operation_fail", OperationFail)
	event.Register("operation_success", OperationSuccess)
}
