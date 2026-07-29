// Package ringer , is where the apps riging happens
package ringer

import (
	"fmt"
	"sync"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

var (
	mu           sync.Mutex
	playing      bool
	streamHandle beep.StreamSeekCloser
	initSpeaker  sync.Once
)

func Start() {
	mu.Lock()
	if playing {
		mu.Unlock()
		return
	}
	playing = true
	mu.Unlock()

	f, err := Assets.Open("assets/alaram.mp3")
	if err != nil {
		fmt.Println("could not open alarm audio:", err)
		mu.Lock()
		playing = false
		mu.Unlock()
		return
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println("could not decode alarm audio:", err)
		mu.Lock()
		playing = false
		mu.Unlock()
		return
	}

	initSpeaker.Do(func() {
		if err := speaker.Init(format.SampleRate, format.SampleRate.N(250_000)); err != nil {
			fmt.Println("could not initialize speaker:", err)
		}
	})

	mu.Lock()
	streamHandle = streamer
	mu.Unlock()

	speaker.Play(beep.Loop(-1, streamer))
}

func Stop() {
	mu.Lock()
	if !playing {
		mu.Unlock()
		return
	}
	playing = false
	streamer := streamHandle
	streamHandle = nil
	mu.Unlock()

	speaker.Clear()
	if streamer != nil {
		if err := streamer.Close(); err != nil {
			fmt.Println("could not close alarm audio:", err)
		}
	}
}
