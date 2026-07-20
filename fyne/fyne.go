package fyne

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/swordde/unsonzer.git/alaram"
)

type Manager12 struct {
	alaram.Manager
}

func (c *Manager12) Thething() {
	h := alaram.Alaram{}

	a := app.NewWithID("com.swordde.unsonzer")

	w := a.NewWindow("Hello World")

	hourEntry := widget.NewEntry()
	minute := widget.NewEntry()
	hourEntry.SetPlaceHolder("Hour")
	minute.SetPlaceHolder("minute")
	createButton := widget.NewButton("Create Alarm", func() {
		h.Hour, _ = strconv.Atoi(hourEntry.Text)
		h.Minute, _ = strconv.Atoi(minute.Text)
		c.Createalaram(h)
	})
	w.SetContent(
		container.NewVBox(
			hourEntry,
			minute,
			createButton,
		),
	)
	w.ShowAndRun()
}
