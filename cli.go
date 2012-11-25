package main

import (
	termbox "github.com/nsf/termbox-go"
	"time"
)

func Client() (err error) {
	err = termbox.Init()
	if err != nil {
		return
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	setString(0, 0, "ParkBench", termbox.AttrBold, termbox.ColorDefault)
	err = termbox.Flush()
	if err != nil {
		return
	}
	time.Sleep(time.Second)
	return
}

func setString(x, y int, s string, fg, bg termbox.Attribute) {
	for c := range s {
		termbox.SetCell(x, y, rune(s[c]), fg, bg)
		x++
	}
}
