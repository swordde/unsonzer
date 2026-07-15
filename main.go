package main

import (
	"fmt"
)

type Alaram struct {
	Hour   int
	minute int
}
type AlaramManeger interface {
	Createalaram()
}

func main() {
	H := Alaram{Hour: 10, minute: 10}

	var c AlaramManeger
	c = &H
	c.Createalaram()
	fmt.Println("hello world!!!")
}

func (c Alaram) Createalaram() {
	fmt.Printf("alaram is created at %v hours and %v minute", c.Hour, c.minute)
}
