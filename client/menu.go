package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/fabian4/kavicat/data"
	"log"
	"strings"
)

func newConnectionForRedis() {

	host := widget.NewEntry()
	host.PlaceHolder = "Input your Redis host here."
	host.SetText("127.0.0.1")
	host.Validator = validation.NewRegexp(
		"((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}",
		"Host Invalid")

	port := widget.NewEntry()
	port.PlaceHolder = "Input your Redis port here."
	port.SetText("6379")
	port.Validator = validation.NewRegexp(
		"^([0-9]|[1-9]\\d|[1-9]\\d{2}|[1-9]\\d{3}|[1-5]\\d{4}|6[0-4]\\d{3}|65[0-4]\\d{2}|655[0-2]\\d|6553[0-5])$",
		"Port Invalid")

	auth := widget.NewPasswordEntry()
	auth.PlaceHolder = "Input your auth code if necessary."

	items := []*widget.FormItem{ // we can specify items in the constructor
		{Text: "Host", Widget: host},
		{Text: "Port", Widget: port},
		{Text: "Auth", Widget: auth},
	}

	form := dialog.NewForm(
		"connection for Redis",
		"connect",
		"cancel",
		items,
		func(bool bool) {
			if bool {
				data.NewRedisConn(host.Text, port.Text, auth.Text)
			}
		},
		GetWindow(),
	)
	form.Resize(fyne.NewSize(400, 250))
	form.Show()
}

func newConnectionForLevelDB() {
	folder := dialog.NewFolderOpen(
		func(uri fyne.ListableURI, err error) {
			if uri == nil {
				log.Println("cancelled")
				return
			}
			data.NewLevelDBConn(strings.ReplaceAll(uri.Path(), "/", "\\\\"))
		}, GetWindow())

	folder.Resize(fyne.NewSize(1000, 800))
	folder.Show()
}

func newConnectionForBadger() {
	// todo: to be done
}
