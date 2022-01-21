package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/fabian4/kavicat/myTheme"
	"log"
)

func New() {
	a := app.New()
	a.Settings().SetTheme(myTheme.NewTheme())
	w := a.NewWindow("kavicat")
	w.Resize(fyne.NewSize(1000, 600))
	w.CenterOnScreen()

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	left := canvas.NewText("left", theme.ForegroundColor())
	middle := canvas.NewText("content", theme.ForegroundColor())
	content := container.New(layout.NewBorderLayout(toolbar, nil, left, nil),
		toolbar, left, middle)

	w.SetContent(content)

	w.ShowAndRun()
}
