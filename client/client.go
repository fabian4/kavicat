package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func New() {
	a := app.NewWithID("github.com/fabian4/kavicat")
	//a.Settings().SetTheme(myTheme.NewTheme())
	w := a.NewWindow("kavicat")
	w.Resize(fyne.NewSize(1000, 600))
	w.CenterOnScreen()

	//w.SetMainMenu(
	//	fyne.NewMainMenu(
	//		fyne.NewMenu("one"),
	//		fyne.NewMenu("two"),
	//		fyne.NewMenu("three"),
	//	),
	//	)

	split := container.NewHSplit(NewConnection(), NewWork())
	split.Offset = 0.15

	content := container.NewBorder(
		container.NewVBox(NewToolBar(), widget.NewSeparator()),
		container.NewVBox(widget.NewSeparator(), NewBottom()),
		nil,
		nil,
		split,
	)

	w.SetContent(content)

	w.ShowAndRun()
}
