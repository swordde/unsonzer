package fyne

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/swordde/unsonzer.git/alaram"
)

func Thething(c *alaram.Manager) {
	h := alaram.Alaram{}

	a := app.NewWithID("com.swordde.unsonzer")

	w := a.NewWindow("Hello World")
	message := widget.NewLabel("welcome")

	clock := widget.NewButton("Button", func() {
		formatted := time.Now().Format("10:20:30")
		message.SetText(formatted)
	})

	hourEntry := widget.NewEntry()
	minute := widget.NewEntry()
	hourEntry.SetPlaceHolder("Hour")
	minute.SetPlaceHolder("minute")
	createButton := widget.NewButton("Create Alarm", func() {
		h.Hour, _ = strconv.Atoi(hourEntry.Text)
		h.Minute, _ = strconv.Atoi(minute.Text)
		h.Trigered = false
		c.Createalaram(h)
		fmt.Println("this is the time:", c)
	})
	w.SetContent(
		container.NewVBox(
			hourEntry,
			minute,
			createButton,
			message,
			clock,
		),
	)
	w.ShowAndRun()
}
