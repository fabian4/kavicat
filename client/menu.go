package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/fabian4/kavicat/conn"
	"github.com/fabian4/kavicat/data"
)

func NewMenu() *fyne.MainMenu {

	return fyne.NewMainMenu(
		newHome(),
		newConn(),
	)
}

func newHome() *fyne.Menu {
	aa := fyne.NewMenuItem("aa", func() {
		success := dialog.NewInformation("Success", "success", GetWindow())
		success.Resize(fyne.NewSize(200, 150))
		success.Show()
	})
	aa.IsQuit = true
	return fyne.NewMenu(
		"home",
		aa,
	)
}

func newConn() *fyne.Menu {

	levelDB := fyne.NewMenuItem("For LevelDB", func() {})

	badger := fyne.NewMenuItem("For Badger", func() {})

	return fyne.NewMenu(
		"connect",
		newConnectionForRedis(),
		levelDB,
		badger,
	)
}

func newConnectionForRedis() *fyne.MenuItem {

	host := widget.NewEntry()
	host.PlaceHolder = "Input your Redis host here."
	host.Validator = validation.NewRegexp(
		"((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}",
		"Host Invalid")

	port := widget.NewEntry()
	port.PlaceHolder = "Input your Redis port here."
	port.Validator = validation.NewRegexp(
		"^([0-9]|[1-9]\\d|[1-9]\\d{2}|[1-9]\\d{3}|[1-5]\\d{4}|6[0-4]\\d{3}|65[0-4]\\d{2}|655[0-2]\\d|6553[0-5])$",
		"Port Invalid")

	auth := widget.NewEntry()
	auth.PlaceHolder = "Input your auth code if necessary."

	name := widget.NewEntry()
	name.PlaceHolder = "Name your connection here."

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
				if bool {
					conn.NewRedisConn(
						&data.RedisConn{
							Host: host.Text,
							Port: port.Text,
							Auth: auth.Text,
							Name: name.Text,
						})
				}
			},
			GetWindow(),
		)
		form.Resize(fyne.NewSize(400, 300))
		form.Show()
	})

	return redis
}
