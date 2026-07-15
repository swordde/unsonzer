package main

import (
	"fmt"
	"time"
)

type Alaram struct {
	Hour   int
	minute int
}
type Manager struct {
	alram []Alaram
}
type AlaramManeger interface {
	Createalaram(a Alaram)
	Schedular()
}

func main() {
	H := Alaram{Hour: 10, minute: 10}

	c := Manager{}
	c.Createalaram(H)
	fmt.Println("Alarams:", c)
	go c.Schedular()
}

func (c *Manager) Createalaram(a Alaram) {
	c.alram = append(c.alram, a)
	fmt.Printf("alaram is created at %v hours and %v minute", a.Hour, a.minute)
}

func (c Manager) Schedular() {
	now := time.Now()
	hour := now.Hour()
	minute := now.Minute()
	for {
		d := c.alram[0]
		if hour == d.Hour && minute == d.Hour {
			fmt.Printf("The alram is going on pls ")
		}
	}
}
