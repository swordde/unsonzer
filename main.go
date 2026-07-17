package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
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
	n := 2
	c := Manager{}
	//	fmt.Scanf("number of alrams:%d", &n)
	for range n {
		H := Alaram{}
		fmt.Scanf("%d", &H.Hour)
		fmt.Scanf("%d", &H.minute)

		c.Createalaram(H)
	}
	fmt.Println("Alarams:", c)
	c.Schedular()
}

func (c *Manager) Createalaram(a Alaram) {
	c.alram = append(c.alram, a)
	fmt.Printf("alaram is created at %v hours and %v minute", a.Hour, a.minute)
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
		if hour == d.Hour && minute == d.minute {
			fmt.Println("The alram is going on pls finish the task to stop the alram")
			Ringer()
			time.Sleep(2 * time.Second)
			j = +1
		} else {
			i = i + 1
			fmt.Printf("not %d yet!!", i)
		}
	}
}

func Ringer() {
	f, _ := os.Open("/home/swordemon/Music/I Need More Space - Jeremy Black.mp3")
	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/100))
	speaker.Play(streamer)
	for {

		fmt.Println("its alarming!!!")
		time.Sleep(10 * time.Second)
		speaker.Close()
		break

	}
}
