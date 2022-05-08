package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/fabian4/kavicat/data"
	"github.com/fabian4/kavicat/event"
	"log"
)

func NewBadgerWork() {
	split := container.NewHSplit(newBadgerKeys(), newBadgerDetail())
	split.Offset = 0.2
	work := container.NewBorder(newBadgerHeadInfo(), nil, nil, nil, split)
	win := GetWindow()
	win.SetContent(work)
}

func newBadgerHeadInfo() fyne.CanvasObject {

	return container.NewBorder(
		nil,
		widget.NewSeparator(),
		widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
			data.CloseBadgerConnection()
			event.Emit("switchUI", "Home")
		}),
		nil,
		widget.NewLabel(data.BadgerConnName),
	)
}

func newBadgerKeys() fyne.CanvasObject {
	bindData := data.BadgerKeys

	list := widget.NewListWithData(bindData,
		func() fyne.CanvasObject {
			return container.NewBorder(nil, nil, nil, widget.NewIcon(theme.MoreVerticalIcon()),
				widget.NewLabel("template"))
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).Bind(i.(binding.String))
		})
	list.OnSelected = func(id widget.ListItemID) {
		data.SetBadgerValuesByKeyId(id)
	}

	return list
}

func newBadgerDetail() fyne.CanvasObject {
	keyLabel := widget.NewLabel("Key: ")
	key := widget.NewEntryWithData(data.BadgerKey)
	key.Wrapping = fyne.TextWrapOff
	key.Validator = nil
	key.SetPlaceHolder("  ")
	key.Disable()

	value := widget.NewEntryWithData(data.BadgerValue)
	value.Wrapping = fyne.TextWrapOff
	value.Validator = nil

	refreshButton := widget.NewButtonWithIcon("", theme.MediaReplayIcon(), func() {
		data.SetBadgerValuesByKey(key.Text)
	})
	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		data.DeleteBadgerValuesByKey(key.Text)
	})
	saveButton := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		data.SaveBadgerValuesByKeyAndValue(key.Text, value.Text)
	})

	addButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		addNewBadgerContent()
	})

	top := container.NewBorder(
		nil,
		nil,
		nil,
		container.NewHBox(refreshButton, deleteButton, saveButton, addButton),
		container.NewHBox(keyLabel, key),
	)

	return container.NewBorder(top, nil, nil, nil, value)
}

func addNewBadgerContent() {
	key := widget.NewEntry()
	key.PlaceHolder = "Input your key here."
	key.Validator = nil

	value := widget.NewEntry()
	value.PlaceHolder = "Input your value here."
	value.Validator = nil

	items := []*widget.FormItem{
		{Text: "key", Widget: key},
		{Text: "value", Widget: value},
	}

	form := dialog.NewForm(
		"Add key and value",
		"add",
		"cancel",
		items,
		func(bool bool) {
			data.SaveBadgerValuesByKeyAndValue(key.Text, value.Text)
			log.Println("save " + key.Text + ": " + value.Text)
		},
		GetWindow(),
	)
	form.Resize(fyne.NewSize(400, 200))
	form.Show()
}
