package fyne

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/swordde/unsonzer.git/alaram"
)

func Thething() {
	//h:=alaram.Alaram{}
	a := app.NewWithID("github.com/swordde/unsonzer.git")

	w := a.NewWindow("Hello World")

	hourEntry := widget.NewEntry()
	minute := widget.NewEntry()
	hourEntry.SetPlaceHolder("Hour")
	minute.SetPlaceHolder("minute")
	createButton := widget.NewButton("Create Alarm", func() {
		fmt.Println(hourEntry.Text)
		fmt.Println(minute.Text)
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
