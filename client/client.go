package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/fabian4/kavicat/myTheme"
)

func New() {
	a := app.New()
	a.Settings().SetTheme(myTheme.NewTheme())
	w := a.NewWindow("kavicat")
	w.Resize(fyne.NewSize(1000, 600))
	w.CenterOnScreen()

	content := container.NewBorder(
		container.NewVBox(NewToolBar(), widget.NewSeparator()),
		container.NewVBox(widget.NewSeparator(), NewBottom()),
		container.NewHBox(NewConnection(), widget.NewSeparator()),
		NewWork())

	w.SetContent(content)

	w.ShowAndRun()
}
