package fyne

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/swordde/unsonzer.git/alaram"
)

func timeupdater(clock *widget.Label) {
	format := time.Now().Format("15:04:05")
	clock.SetText(format)
}

func Thething(c *alaram.Manager) {
	h := alaram.Alaram{}

	a := app.NewWithID("com.swordde.unsonzer")

	w := a.NewWindow("Hello World")

	clock := widget.NewLabel("")

	timeupdater(clock)

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
			clock,
		),
	)
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			fyne.Do(func() {
				timeupdater(clock)
			})
		}
	}()

	w.ShowAndRun()
}
