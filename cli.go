package main

import (
	t "github.com/nsf/termbox-go"
	"strings"
)

var (
	Queue = make(chan t.Event)
)

//
func newMessage(s string) {

}

func setString(x, y int, s string, fg, bg t.Attribute) {
	for c := range s {
		t.SetCell(x, y, rune(s[c]), fg, bg)
		x++
	}
}

//interpret takes a string argument as a command or message typed on the input line. It interprets it and, using global variables, performs the appropriate actions.
func interpret(input string) {
	if strings.HasPrefix(input, "/") {
		//If the input is a command, then
		//switch on everything after "/"
		switch input[1:] {
		case "quit":
			close(Queue)
		}
	}
}

func loopIn(queue <-chan t.Event) (err error) {
	var input string
	var x, y int
	_, y = t.Size()
	y--

	for ev := range queue {
		switch ev.Type {
		case t.EventKey:
			switch ev.Key {
			case t.KeyEsc:
				//If the key pressed is esc,
				//then return immediately.
				return
			case t.KeyEnter:
				//If the user presses enter,
				//interpret the input, and
				//possibly send a message.
				interpret(input)
			default:
				input += string(ev.Ch)
				t.SetCell(x, y, ev.Ch, t.ColorDefault, t.ColorDefault)
				err = t.Flush()
				if err != nil {
					return
				}
				x++
			}
		}
	}
	//If we get here, the channel has closed.
	return
}
