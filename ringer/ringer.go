// Package ringer , is where the apps riging happens
package ringer

import (
	"fmt"
	"os"
	"time"

	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func Ringer() {
	f, err := os.Open("/home/swordemon/Music/I Need More Space - Jeremy Black.mp3")
	if err != nil {
		panic("err")
	}
	streamer, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/100))
	speaker.Play(streamer)
	for {

		fmt.Println("its alarming!!!")
		time.Sleep(10 * time.Second)
		speaker.Clear()
		break

	}
}
