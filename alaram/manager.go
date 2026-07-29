// Package alaram this is where the alram things liesss.
package alaram

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/swordde/unsonzer.git/ringer"
)

type Alarm struct {
	Hour              int
	Minute            int
	Triggered         bool
	Enabled           bool
	Difficulty        string
	ChallengeQuestion string
	ChallengeAnswer   string
}

type ActiveChallenge struct {
	Alarm      Alarm
	Question   string
	Answer     string
	Difficulty string
}

type Manager struct {
	alarms []Alarm
	mu     sync.RWMutex

	activeChallenge *ActiveChallenge
}

type AlaramManeger interface {
	Createalaram(a Alarm)
	Schedular()
}

func (c *Manager) Createalaram(a Alarm) *Manager {
	c.mu.Lock()
	defer c.mu.Unlock()

	a.Enabled = true
	a.Triggered = false
	a.Difficulty = strings.ToLower(strings.TrimSpace(a.Difficulty))
	if a.Difficulty == "" {
		a.Difficulty = "easy"
	}
	question, answer := buildChallenge(a.Difficulty)
	a.ChallengeQuestion = question
	a.ChallengeAnswer = answer

	c.alarms = append(c.alarms, a)
	fmt.Printf("alarm created at %02d:%02d with %s challenge\n", a.Hour, a.Minute, a.Difficulty)
	return c
}

func (c *Manager) Alarms() []Alarm {
	c.mu.RLock()
	defer c.mu.RUnlock()

	out := make([]Alarm, len(c.alarms))
	copy(out, c.alarms)
	return out
}

func (c *Manager) ActiveChallenge() (ActiveChallenge, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.activeChallenge == nil {
		return ActiveChallenge{}, false
	}

	return *c.activeChallenge, true
}

func (c *Manager) SubmitAnswer(answer string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.activeChallenge == nil {
		return false
	}

	if strings.EqualFold(strings.TrimSpace(answer), strings.TrimSpace(c.activeChallenge.Answer)) {
		c.activeChallenge.Alarm.Triggered = false
		c.activeChallenge.Alarm.Enabled = true
		c.activeChallenge = nil
		ringer.Stop()
		return true
	}

	return false
}

func (c *Manager) Schedular() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()

		for i := range c.alarms {
			alarm := &c.alarms[i]
			if alarm.Triggered && (now.Hour() != alarm.Hour || now.Minute() != alarm.Minute) {
				alarm.Triggered = false
			}

			if alarm.Enabled && !alarm.Triggered && now.Hour() == alarm.Hour && now.Minute() == alarm.Minute {
				alarm.Triggered = true
				c.activeChallenge = &ActiveChallenge{
					Alarm:      *alarm,
					Question:   alarm.ChallengeQuestion,
					Answer:     alarm.ChallengeAnswer,
					Difficulty: alarm.Difficulty,
				}
				fmt.Printf("alarm triggered: %02d:%02d\n", alarm.Hour, alarm.Minute)
				ringer.Start()
			}
		}

		c.mu.Unlock()
	}
}

func buildChallenge(difficulty string) (string, string) {
	switch difficulty {
	case "medium":
		n1 := rand.Intn(20) + 5
		n2 := rand.Intn(20) + 5
		question := fmt.Sprintf("Medium mode: %d - %d = ?", n1, n2)
		return question, strconv.Itoa(n1 - n2)
	case "hard":
		n1 := rand.Intn(12) + 3
		n2 := rand.Intn(12) + 3
		question := fmt.Sprintf("Hard mode: %d × %d = ?", n1, n2)
		return question, strconv.Itoa(n1 * n2)
	default:
		n1 := rand.Intn(10) + 1
		n2 := rand.Intn(10) + 1
		question := fmt.Sprintf("Easy mode: %d + %d = ?", n1, n2)
		return question, strconv.Itoa(n1 + n2)
	}
}
