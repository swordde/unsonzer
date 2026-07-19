package main

import (
	"fmt"

	"github.com/swordde/unsonzer.git/fyne"

	"github.com/swordde/unsonzer.git/alaram"
)

func main() {
	fyne.Thething()
	n := 2

	c := alaram.Manager{}
	//	fmt.Scanf("number of alrams:%d", &n)
	for range n {
		H := alaram.Alaram{}
		fmt.Scanf("%d", &H.Hour)
		fmt.Scanf("%d", &H.Minute)

		c.Createalaram(H)
	}
	fmt.Println("Alarams:", c)
	go c.Schedular()
}
