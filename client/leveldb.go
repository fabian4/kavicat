package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/fabian4/kavicat/data"
	"github.com/fabian4/kavicat/event"
)

func NewLevelDBWork() {
	split := container.NewHSplit(newLevelDBKeys(), newLevelDBDetail())
	split.Offset = 0.2
	work := container.NewBorder(newLevelDBHeadInfo(), nil, nil, nil, split)
	win := GetWindow()
	win.SetContent(work)
}

func newLevelDBHeadInfo() fyne.CanvasObject {

	return container.NewBorder(
		nil,
		widget.NewSeparator(),
		widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
			event.Emit("switchUI", "Home")
		}),
		nil,
		widget.NewLabel(data.LevelDBConnName),
	)
}

func newLevelDBKeys() fyne.CanvasObject {
	bindData := data.LevelDBKeys

	list := widget.NewListWithData(bindData,
		func() fyne.CanvasObject {
			return container.NewBorder(nil, nil, nil, widget.NewIcon(theme.MoreVerticalIcon()),
				widget.NewLabel("template"))
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).Bind(i.(binding.String))
		})
	list.OnSelected = func(id widget.ListItemID) {
		data.SetLevelDBValuesByKeyId(id)
	}

	return list
}

func newLevelDBDetail() fyne.CanvasObject {
	keyLabel := widget.NewLabel("Key")
	key := widget.NewEntryWithData(data.LevelDBKey)
	key.Wrapping = fyne.TextWrapOff
	key.Validator = nil
	key.SetPlaceHolder("  ")
	key.Disable()
	category := widget.NewButton("String", func() {})

	category.Disable()

	TtlLabel := widget.NewLabel("TTL")
	TTL := widget.NewEntry()
	TTL.SetText("-1")

	value := widget.NewEntryWithData(data.LevelDBValue)
	value.Validator = nil

	//refreshButton := widget.NewButtonWithIcon("", theme.MediaReplayIcon(), func() {
	//	data.SetValuesByKey(key.Text)
	//})
	//deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
	//	data.DeleteValuesByKey(key.Text)
	//})
	//saveButton := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
	//	data.SaveValuesByKeyAndValue(key.Text, value.Text)
	//})

	//addButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
	//	addNewContent()
	//})

	top := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		//container.NewHBox(category, refreshButton, deleteButton, saveButton, addButton),
		container.NewHBox(keyLabel, key, TtlLabel, TTL),
	)

	return container.NewBorder(top, nil, nil, nil, value)
}
