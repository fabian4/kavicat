package client

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func NewWork() fyne.CanvasObject {

	//return canvas.NewText("content", theme.ForegroundColor())

	return container.NewBorder(newHeadInfo(), nil, nil, nil, newKeys())
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
			fmt.Println("selected", s)
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
			//widget.NewLabel("127.0.0.1:6379"),
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
	data := make([]string, 8)
	for i := range data {
		data[i] = "Test Item " + strconv.Itoa(i)
	}

	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")
	center := container.NewHBox(icon, label)

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.StorageIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(data[id])
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		label.SetText(data[id])
		icon.SetResource(theme.DocumentIcon())
	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}
	//list.Select(125)

	split := container.NewHSplit(list, container.NewCenter(center))
	split.Offset = 0.2

	return split
}
