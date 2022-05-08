package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/fabian4/kavicat/event"
	"github.com/fabian4/kavicat/resource"
)

var (
	App    fyne.App
	Window fyne.Window
)

func Init() {

	RegisterEvents()

	App = app.NewWithID("github.com/fabian4/kavicat")
	Window = App.NewWindow("kavicat")
	Window.Resize(fyne.NewSize(1000, 600))
	Window.CenterOnScreen()
	NewHome()

	Window.ShowAndRun()
}

func NewHome() {
	img := canvas.NewImageFromResource(resource.ResourceLogoPng)
	img.FillMode = canvas.ImageFillOriginal
	redisButton := widget.NewButton("Establish connection for Redis", newConnectionForRedis)
	badgerButton := widget.NewButton("Establish connection for Badger", newConnectionForBadger)
	levelDBButton := widget.NewButton("Establish connection for LevelDB", newConnectionForLevelDB)
	buttons := container.NewVBox(redisButton, badgerButton, levelDBButton)
	Window.SetContent(container.NewVBox(img, container.NewCenter(buttons)))
}

func RegisterEvents() {
	event.Register("empty", Empty)
	event.Register("switchUI", SwitchUI)
	event.Register("operation_fail", OperationFail)
	event.Register("connection_fail", ConnectionFail)
	event.Register("operation_success", OperationSuccess)
	event.Register("connection_success", ConnectionSuccess)
}
