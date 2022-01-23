package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func NewToolBar(client *Client) fyne.CanvasObject {

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			dialog.ShowForm("Establish a new connection", "add", "aha", newConnection(), func(bool bool) {}, client.Window)
			//dialog.ShowCustom("Establish a new connection", "aha", container.NewMax(newConnection()), client.Window)
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	return toolbar
}

func newConnection() []*widget.FormItem {

	//entry := widget.NewEntry()
	//textArea := widget.NewMultiLineEntry()

	return []*widget.FormItem{ // we can specify items in the constructor
		{Text: "Entry", Widget: widget.NewEntry()},
		{Text: "Entry", Widget: widget.NewEntry()},
		{Text: "Entry", Widget: widget.NewEntry()},
		{Text: "Entry", Widget: widget.NewEntry()},
	}
}
