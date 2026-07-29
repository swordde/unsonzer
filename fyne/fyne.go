package fyne

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/swordde/unsonzer.git/alaram"
)

func timeupdater(clock *canvas.Text) {
	format := time.Now().Format("15:04:05")
	clock.Text = format
}

func nextAlarmText(alarms []alaram.Alarm) string {
	if len(alarms) == 0 {
		return "No alarms scheduled"
	}

	var next *alaram.Alarm
	now := time.Now()
	soonest := 24 * time.Hour

	for i := range alarms {
		alarm := &alarms[i]
		alarmTime := time.Date(now.Year(), now.Month(), now.Day(), alarm.Hour, alarm.Minute, 0, 0, now.Location())
		if alarmTime.Before(now) {
			alarmTime = alarmTime.Add(24 * time.Hour)
		}
		delta := time.Until(alarmTime)
		if delta < soonest {
			soonest = delta
			next = alarm
		}
	}

	if next == nil {
		return "No alarms scheduled"
	}

	return fmt.Sprintf("Next alarm %02d:%02d • %s", next.Hour, next.Minute, titleCase(next.Difficulty))
}

func nextAlarmCountdown(alarms []alaram.Alarm) string {
	if len(alarms) == 0 {
		return "Set your first alarm"
	}

	now := time.Now()
	var best time.Duration = 24 * time.Hour
	for i := range alarms {
		alarm := alarms[i]
		alarmTime := time.Date(now.Year(), now.Month(), now.Day(), alarm.Hour, alarm.Minute, 0, 0, now.Location())
		if alarmTime.Before(now) {
			alarmTime = alarmTime.Add(24 * time.Hour)
		}
		delta := time.Until(alarmTime)
		if delta < best {
			best = delta
		}
	}

	hours := int(best.Hours())
	minutes := int(best.Minutes()) % 60
	if hours <= 0 {
		return fmt.Sprintf("%d min away", minutes)
	}
	return fmt.Sprintf("%dh %dm away", hours, minutes)
}

func difficultySummary(difficulty string) string {
	switch difficulty {
	case "medium":
		return "2-3 challenges"
	case "hard":
		return "3+ challenges"
	default:
		return "1 challenge"
	}
}

func titleCase(input string) string {
	if input == "" {
		return ""
	}
	return strings.ToUpper(input[:1]) + input[1:]
}

func Thething(c *alaram.Manager) {
	a := app.NewWithID("com.swordde.unsonzer")
	w := a.NewWindow("Unsonzer")
	w.Resize(fyne.NewSize(430, 860))
	a.Settings().SetTheme(theme.DarkTheme())

	clockText := canvas.NewText("", theme.ForegroundColor())
	clockText.Alignment = fyne.TextAlignCenter
	clockText.TextStyle = fyne.TextStyle{Bold: true}
	clockText.TextSize = 48
	timeupdater(clockText)

	subtitle := widget.NewLabel("Rise up and build a wake-up routine that actually works.")
	subtitle.Wrapping = fyne.TextWrapWord
	subtitle.Alignment = fyne.TextAlignCenter

	headerTitle := canvas.NewText("Good morning", theme.ForegroundColor())
	headerTitle.Alignment = fyne.TextAlignCenter
	headerTitle.TextStyle = fyne.TextStyle{Bold: true}
	headerTitle.TextSize = 18

	statusLabel := widget.NewLabel("Set a time, choose a challenge, and build your wake-up ritual.")
	statusLabel.Wrapping = fyne.TextWrapWord
	statusLabel.Alignment = fyne.TextAlignCenter

	alarms := c.Alarms()
	nextHero := widget.NewCard("Next Alarm", nextAlarmText(alarms), container.NewVBox(
		clockText,
		widget.NewLabelWithStyle(nextAlarmCountdown(alarms), fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		statusLabel,
	))
	nextHero.SetSubTitle("A minimal, focused alarm screen")

	hourEntry := widget.NewEntry()
	minuteEntry := widget.NewEntry()
	hourEntry.SetPlaceHolder("Hour (0-23)")
	minuteEntry.SetPlaceHolder("Minute (0-59)")

	selectedDifficulty := "easy"
	difficultySelect := widget.NewSelect([]string{"easy", "medium", "hard"}, func(value string) {
		selectedDifficulty = value
	})
	difficultySelect.SetSelected("easy")
	difficultyHint := widget.NewLabel("Wake mode: " + difficultySummary(selectedDifficulty))
	difficultyHint.Alignment = fyne.TextAlignCenter

	alarmList := container.NewVBox()
	challengeQuestion := widget.NewLabel("")
	challengeQuestion.Wrapping = fyne.TextWrapWord
	challengeQuestion.Alignment = fyne.TextAlignCenter
	challengeEntry := widget.NewEntry()
	challengeEntry.SetPlaceHolder("Type your answer")
	var refreshAlarmList func()
	var updateChallengeCard func()
	var challengeTitle *widget.Card
	var challengeBody *fyne.Container
	var statsRow *fyne.Container

	makeStatCard := func(title, value, subtitle string) *widget.Card {
		return widget.NewCard(title, subtitle, container.NewCenter(widget.NewLabelWithStyle(value, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})))
	}

	refreshAlarmList = func() {
		alarms := c.Alarms()
		rows := make([]fyne.CanvasObject, 0, len(alarms)+1)
		if len(alarms) == 0 {
			rows = append(rows, widget.NewCard("No alarms yet", "Add your first wake-up", container.NewVBox(widget.NewLabel("Create one above and it will appear here."))))
		} else {
			for _, alarm := range alarms {
				state := "armed"
				if alarm.Triggered {
					state = "ringing"
				}

				challengeCopy := difficultySummary(alarm.Difficulty)
				alarmBody := container.NewHBox(
					widget.NewLabelWithStyle(fmt.Sprintf("%02d:%02d", alarm.Hour, alarm.Minute), fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
					widget.NewLabel(state),
				)
				rows = append(rows, widget.NewCard("Alarm", fmt.Sprintf("%s • %s", titleCase(alarm.Difficulty), challengeCopy), alarmBody))
			}
		}
		alarmList.Objects = rows
		alarmList.Refresh()
	}

	updateStats := func() {
		alarms = c.Alarms()
		nextHero.SetTitle(nextAlarmText(alarms))
		nextHero.SetSubTitle(nextAlarmCountdown(alarms))
		statsRow.Objects = []fyne.CanvasObject{
			makeStatCard("Wake streak", fmt.Sprintf("%d", len(alarms)+18), "Days"),
			makeStatCard("Wake score", fmt.Sprintf("%d", 90-len(alarms)*2), "Excellent"),
		}
		statsRow.Refresh()
	}

	updateChallengeCard = func() {
		active, ok := c.ActiveChallenge()
		if !ok {
			challengeTitle.Hide()
			challengeBody.Hide()
			return
		}

		challengeQuestion.SetText(active.Question)
		challengeTitle.SetSubTitle(active.Difficulty)
		challengeTitle.Show()
		challengeBody.Show()
	}

	challengeBody = container.NewVBox(
		challengeQuestion,
		challengeEntry,
		widget.NewButton("Solve challenge", func() {
			if c.SubmitAnswer(challengeEntry.Text) {
				challengeEntry.SetText("")
				statusLabel.SetText("Alarm stopped. Nice work.")
				refreshAlarmList()
				updateStats()
				updateChallengeCard()
			} else {
				statusLabel.SetText("Not quite. Try again.")
			}
		}),
	)
	challengeBody.Hide()
	challengeTitle = widget.NewCard("Wake-up challenge", "Solve to stop the alarm", challengeBody)
	challengeTitle.Hide()

	createButton := widget.NewButtonWithIcon("Save Alarm", theme.ContentAddIcon(), func() {
		hour, err := strconv.Atoi(hourEntry.Text)
		if err != nil || hour < 0 || hour > 23 {
			statusLabel.SetText("Choose an hour from 0 to 23.")
			return
		}

		minute, err := strconv.Atoi(minuteEntry.Text)
		if err != nil || minute < 0 || minute > 59 {
			statusLabel.SetText("Choose a minute from 0 to 59.")
			return
		}

		alarm := alaram.Alarm{
			Hour:       hour,
			Minute:     minute,
			Difficulty: selectedDifficulty,
		}
		c.Createalaram(alarm)
		statusLabel.SetText(fmt.Sprintf("Alarm set for %02d:%02d with a %s challenge.", hour, minute, selectedDifficulty))
		hourEntry.SetText("")
		minuteEntry.SetText("")
		difficultySelect.SetSelected("easy")
		selectedDifficulty = "easy"
		difficultyHint.SetText("Wake mode: " + difficultySummary(selectedDifficulty))
		refreshAlarmList()
		updateStats()
	})

	formCard := widget.NewCard("Create alarm", "Quick setup", container.NewVBox(
		hourEntry,
		minuteEntry,
		difficultySelect,
		difficultyHint,
		createButton,
	))

	statsRow = container.NewGridWithColumns(2,
		makeStatCard("Wake streak", "18", "Days"),
		makeStatCard("Wake score", "92", "Excellent"),
	)

	weeklyCard := widget.NewCard("Weekly performance", "Simple overview", container.NewVBox(
		widget.NewLabel("Mon   Tue   Wed   Thu   Fri   Sat   Sun"),
		widget.NewProgressBar(),
	))

	body := container.NewVBox(
		headerTitle,
		subtitle,
		nextHero,
		statsRow,
		weeklyCard,
		formCard,
		challengeTitle,
		widget.NewLabelWithStyle("Scheduled alarms", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		alarmList,
	)

	refreshAlarmList()
	updateStats()
	updateChallengeCard()

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			fyne.Do(func() {
				timeupdater(clockText)
				clockText.Refresh()
				refreshAlarmList()
				updateChallengeCard()
			})
		}
	}()

	w.SetContent(container.NewVScroll(container.NewPadded(body)))
	w.ShowAndRun()
}
