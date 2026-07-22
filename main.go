package main

import (
	"fmt"

	"github.com/swordde/unsonzer.git/alaram"
	"github.com/swordde/unsonzer.git/fyne"
)

func main() {
	c := alaram.Manager{}

	go c.Schedular()
	fyne.Thething(&c)

	//	fmt.Scanf("number of alrams:%d", &n)

	fmt.Println("Alarams:", c)
}
