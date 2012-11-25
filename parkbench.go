package main

import (
	//"github.com/SashaCrofter/benchgolib"
	t "github.com/nsf/termbox-go"
	"log"
)

func main() {
	err := start()
	if err != nil {
		log.Println(err)
		println("An error was encountered.")
	}
	//If error wasn't nil, then we're exiting
	//gracefully.
}

func start() (err error) {
	err = t.Init()
	if err != nil {
		return
	}
	defer t.Close()
	_, y := t.Size()
	t.Clear(Fg, Bg)

	setString(0, 0, "ParkBench", Fg|t.AttrBold, Bg)
	setString(0, y-2, "Use '/connect <ipv6>' to chat with a friend.", Fg, Bg)

	err = t.Flush()
	if err != nil {
		return
	}

	//Queue is declared as a global variable
	//in cli.go.
	go func(queue chan<- t.Event) {
		for {
			queue <- t.PollEvent()
		}
	}(Queue)

	//Now, just take user input until the user exits.
	err = loopIn("> ", Queue)
	return
}
