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
	H := Alaram{}
	fmt.Scanf("%d", &H.Hour)
	fmt.Scanf("%d", &H.minute)

	c := Manager{}
	c.Createalaram(H)
	fmt.Println("Alarams:", c)
	c.Schedular()
}

func (c *Manager) Createalaram(a Alaram) {
	c.alram = append(c.alram, a)
	fmt.Printf("alaram is created at %v hours and %v minute", a.Hour, a.minute)
}

func (c Manager) Schedular() {
	now := time.Now()
	hour := now.Hour()

	minute := now.Minute()
	fmt.Printf("%v", hour)
	for {
		d := c.alram[0]
		if hour == d.Hour && minute == d.minute {
			fmt.Printf("The alram is going on pls finish the task to stop the alram")
			Ringer()
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("exited")
			break
		}
	}
}

func Ringer() {
	for {

		fmt.Println("its alarming!!!")

		time.Sleep(1 * time.Minute)
		break

	}
}
