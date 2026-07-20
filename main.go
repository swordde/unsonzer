package main

import (
	"fmt"

	"github.com/swordde/unsonzer.git/alaram"
	"github.com/swordde/unsonzer.git/fyne"
)

func main() {
	k := &fyne.Manager12{}
	c := alaram.Manager{}

	go c.Schedular()
	k.Thething()

	//	fmt.Scanf("number of alrams:%d", &n)

	fmt.Println("Alarams:", c)
}
