// Package ringer , is where the apps riging happens
package ringer

import (
	"fmt"
	"time"

	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func Ringer() {
	f, err := Assets.Open("assets/alaram.mp3")
	if err != nil {
		panic(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		panic(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second))
	speaker.Play(streamer)
	for {

		fmt.Println("its alarming!!!")
		time.Sleep(10 * time.Second)
		speaker.Clear()
		break

	}
}
