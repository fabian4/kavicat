package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewMenu(client *Client) *fyne.MainMenu {
	return fyne.NewMainMenu(
		newHome(),
		newConn(client),
	)
}

func newHome() *fyne.Menu {
	aa := fyne.NewMenuItem("aa", func() {})
	aa.IsQuit = true
	return fyne.NewMenu(
		"home",
		aa,
	)
}

func newConn(client *Client) *fyne.Menu {

	levelDB := fyne.NewMenuItem("For LevelDB", func() {})

	badger := fyne.NewMenuItem("For Badger", func() {})

	return fyne.NewMenu(
		"connect",
		newConnectionForRedis(client),
		levelDB,
		badger,
	)
}

func newConnectionForRedis(client *Client) *fyne.MenuItem {

	host := widget.NewEntry()
	host.Validator = validation
	port := widget.NewEntry()
	auth := widget.NewEntry()
	name := widget.NewEntry()

	items := []*widget.FormItem{ // we can specify items in the constructor
		{Text: "Host", Widget: host},
		{Text: "Port", Widget: port},
		{Text: "Auth", Widget: auth},
		{Text: "Name", Widget: name},
	}

	redis := fyne.NewMenuItem("For Redis", func() {
		form := dialog.NewForm(
			"connection for Redis",
			"connect",
			"cancel",
			items,
			func(bool bool) {
				println("connect")
			},
			client.Window,
		)
		form.Resize(fyne.NewSize(400, 300))
		form.Show()
	})

	return redis
}
