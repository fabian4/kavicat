package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/fabian4/kavicat/data"
)

func NewWork() fyne.CanvasObject {
	split := container.NewHSplit(newKeys(), newDetail())
	split.Offset = 0.2
	return container.NewBorder(newHeadInfo(), nil, nil, nil, split)
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
			//fmt.Println("selected", s)
		},
	)
	selectEntry.SetSelectedIndex(0)
	//info := container.NewHBox(
	//	widget.NewLabel("redis: 127.0.0.1:6379"),
	//	widget.NewLabel("db:"),
	//	selectEntry,
	//)
	//return container.NewVBox(info, widget.NewSeparator())

	//rdc := redis.NewClient(&redis.Options{
	//	Addr:     "127.0.0.1:6379",
	//	Password: "",
	//	DB:       0,
	//})
	//
	//result := rdc.Info(context.Background())
	//
	//info, _ := redisinfo.Parse(result.String())
	//
	//data, _ := json.Marshal(info)
	//fmt.Printf("%s\n", data)

	//log.Println(json.Marshal(info))

	return container.NewBorder(
		nil,
		widget.NewSeparator(),
		nil,
		container.NewHBox(
			//widget.NewLabel("clients: "+strconv.FormatInt(info.Clients.ConnectedClients, 10)),
			widget.NewLabel("clients: 3"),
			widget.NewSeparator(),
			//widget.NewLabel("memory: "+info.Memory.UsedMemoryHuman),
			widget.NewLabel("memory: 7349k"),
			widget.NewSeparator(),
			//widget.NewLabel("keys: "+strconv.FormatUint(info.Keyspace[0].Keys, 10)),
			widget.NewLabel("keys: 6"),
			widget.NewSeparator(),
			selectEntry,
		),
		widget.NewLabel("127.0.0.1:6379"),
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
	//key.Refresh()
	category := widget.NewButton("String", func() {})

	category.Disable()

	TtlLabel := widget.NewLabel("TTL")
	TTL := widget.NewEntry()
	TTL.SetText("-1")

	refreshButton := widget.NewButtonWithIcon("", theme.MediaReplayIcon(), func() {

	})
	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {})
	saveButton := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {})

	top := container.NewBorder(
		nil,
		nil,
		nil,
		container.NewHBox(category, refreshButton, deleteButton, saveButton),
		container.NewHBox(keyLabel, key, TtlLabel, TTL),
	)

	value := widget.NewEntryWithData(data.Value)
	value.Validator = nil

	return container.NewBorder(top, nil, nil, nil, value)
}
