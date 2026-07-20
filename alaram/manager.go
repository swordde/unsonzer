// Package alaram this is where the alram  things liesss.
package alaram

import (
	"fmt"
	"time"

	"github.com/swordde/unsonzer.git/ringer"
)

type Alaram struct {
	Hour   int
	Minute int
}
type Manager struct {
	alram []Alaram
}

type AlaramManeger interface {
	Createalaram(a Alaram)
	Schedular()
}

func (c *Manager) Createalaram(a Alaram) *Manager {
	c.alram = append(c.alram, a)
	fmt.Printf("alaram is created at %v hours and %v minute", a.Hour, a.Minute)
	return c
}

func (c *Manager) Schedular() {
	for {
		fmt.Println("schedular")
		now := time.Now()
		fmt.Println(c.alram)

		for _, alarm := range c.alram {
			fmt.Println("for loop")
			if now.Hour() == alarm.Hour &&
				now.Minute() == alarm.Minute {

				fmt.Println("Alarm triggered")
				ringer.Ringer()
			}
		}

		time.Sleep(time.Second)
	}
}
