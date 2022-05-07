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
	"strconv"
)

func NewRedisWork() {
	split := container.NewHSplit(newKeys(), newDetail())
	split.Offset = 0.2
	work := container.NewBorder(newHeadInfo(), nil, nil, nil, split)
	win := GetWindow()
	win.SetContent(work)
}

func newHeadInfo() fyne.CanvasObject {
	selectEntry := widget.NewSelect(
		[]string{
			"DB-0",
			"DB-1",
			"DB-2",
			"DB-3",
			"DB-4",
			"DB-5",
			"DB-6",
			"DB-7",
			"DB-8",
			"DB-9",
			"DB-10",
			"DB-11",
			"DB-12",
			"DB-13",
			"DB-14",
			"DB-15",
		},
		func(s string) {
			index, _ := strconv.Atoi(s[3:])
			data.SwitchDB(index)
		},
	)
	selectEntry.SetSelectedIndex(0)

	return container.NewBorder(
		nil,
		widget.NewSeparator(),
		widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
			event.Emit("switchUI", "Home")
		}),
		container.NewHBox(
			widget.NewLabelWithData(data.Client),
			widget.NewSeparator(),
			widget.NewLabelWithData(data.Memory),
			widget.NewSeparator(),
			widget.NewLabelWithData(data.Count),
			widget.NewSeparator(),
			selectEntry,
		),
		widget.NewLabel(data.GetConnName()),
	)
}

func newKeys() fyne.CanvasObject {
	bindData := data.Keys

	list := widget.NewListWithData(bindData,
		func() fyne.CanvasObject {
			return container.NewBorder(nil, nil, nil, widget.NewIcon(theme.MoreVerticalIcon()),
				widget.NewLabel("template"))
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).Bind(i.(binding.String))
		})
	list.OnSelected = func(id widget.ListItemID) {
		data.SetValuesByKeyId(id)
	}

	return list
}

func newDetail() fyne.CanvasObject {
	keyLabel := widget.NewLabel("Key")
	key := widget.NewEntryWithData(data.Key)
	key.Wrapping = fyne.TextWrapOff
	key.Validator = nil
	key.SetPlaceHolder("  ")
	key.Disable()
	category := widget.NewButton("String", func() {})

	category.Disable()

	TtlLabel := widget.NewLabel("TTL")
	TTL := widget.NewEntry()
	TTL.SetText("-1")

	value := widget.NewEntryWithData(data.Value)
	value.Validator = nil

	refreshButton := widget.NewButtonWithIcon("", theme.MediaReplayIcon(), func() {
		data.SetValuesByKey(key.Text)
	})
	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		data.DeleteValuesByKey(key.Text)
	})
	saveButton := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		data.SaveValuesByKeyAndValue(key.Text, value.Text)
	})

	addButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		addNewContent()
	})

	top := container.NewBorder(
		nil,
		nil,
		nil,
		container.NewHBox(category, refreshButton, deleteButton, saveButton, addButton),
		container.NewHBox(keyLabel, key, TtlLabel, TTL),
	)

	return container.NewBorder(top, nil, nil, nil, value)
}

func addNewContent() {

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
			data.SaveValuesByKeyAndValue(key.Text, value.Text)
			log.Println("save " + key.Text + ": " + value.Text)
		},
		GetWindow(),
	)
	form.Resize(fyne.NewSize(400, 200))
	form.Show()
}
