package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Client struct {
	App    fyne.App
	Window fyne.Window
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Build() {
	a := app.NewWithID("github.com/fabian4/kavicat")
	//a.Settings().SetTheme(customer.NewTheme())
	w := a.NewWindow("kavicat")
	w.Resize(fyne.NewSize(1000, 600))
	w.CenterOnScreen()

	client.App = a
	client.Window = w

	w.SetMainMenu(NewMenu(client))

	//head := container.NewVBox(NewToolBar(client), widget.NewSeparator())

	bottom := container.NewVBox(widget.NewSeparator(), NewBottom())

	split := container.NewHSplit(NewConnection(), NewWork())
	split.Offset = 0.15

	content := container.NewBorder(nil, bottom, nil, nil, split)

	w.SetContent(content)

	w.ShowAndRun()
}
