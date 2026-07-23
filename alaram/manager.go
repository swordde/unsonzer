// Package alaram this is where the alram  things liesss.
package alaram

import (
	"fmt"
	"time"

	"github.com/swordde/unsonzer.git/ringer"
)

type Alaram struct {
	Hour     int
	Minute   int
	Trigered bool
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

		for i := range c.alram {
			alarm := &c.alram[i]
			fmt.Println("for loop")
			fmt.Println("the alrams are:", c.alram)
			if now.Hour() == alarm.Hour &&
				now.Minute() == alarm.Minute {
				fmt.Println(alarm.Trigered)

				if !alarm.Trigered {
					fmt.Println("Alarm triggered")

					alarm.Trigered = true

					ringer.Ringer()
				}
			}
		}

		time.Sleep(time.Second)
	}
}
