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

func (c *Manager) Createalaram(a Alaram) {
	c.alram = append(c.alram, a)
	fmt.Printf("alaram is created at %v hours and %v minute", a.Hour, a.Minute)
}

func (c Manager) Schedular() {
	i := 0
	j := 0
	for {
		now := time.Now()
		hour := now.Hour()

		minute := now.Minute()
		fmt.Printf("%v", hour)

		time.Sleep(1 * time.Second)
		d := c.alram[j]
		if hour == d.Hour && minute == d.Minute {
			fmt.Println("The alram is going on pls finish the task to stop the alram")
			ringer.Ringer()
			time.Sleep(2 * time.Second)
			if len(c.alram)-1 != j {
				j = j + 1
			}

		} else {
			i = i + 1
			fmt.Printf("not %d yet!!", i)

		}
	}
}
