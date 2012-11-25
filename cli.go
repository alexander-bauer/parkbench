package main

import (
	t "github.com/nsf/termbox-go"
	"strings"
)

const (
	Fg = t.ColorDefault //Default foreground color
	Bg = t.ColorDefault //Default background color
)

var (
	Queue = make(chan t.Event)
)

//
func newMessage(s string) {

}

//Uses a simple for loop to write a string of characters on the screen in a single line. It does not flush to the screen.
func setString(x, y int, s string, fg, bg t.Attribute) {
	for c := range s {
		t.SetCell(x, y, rune(s[c]), fg, bg)
		x++
	}
}

//interpret takes a string argument as a command or message typed on the input line. It interprets it and, using global variables, performs the appropriate actions.
func interpret(input string) {
	if strings.HasPrefix(input, "/") {
		input = strings.ToLower(input)
		//If the input is a command, then
		//switch on everything after "/"
		switch input[1:] {
		case "quit":
			close(Queue)
		}
	}
}

func loopIn(prompt string, queue <-chan t.Event) (err error) {
	var input string

	x := len(prompt)
	xMin := x

	xMax, y := t.Size()
	xMax--
	y--

	setString(0, y, prompt, Fg, Bg)
	err = t.Flush()
	if err != nil {
		return
	}

	for ev := range queue {
		switch ev.Type {
		case t.EventKey:
			switch ev.Key {
			case t.KeyEsc, t.KeyCtrlC, t.KeyCtrlD:
				//If the key pressed is esc,
				//then return immediately.
				return
			case t.KeyEnter:
				//If the user presses enter,
				//interpret the input, and
				//possibly send a message.
				interpret(input)

				//Now blank the buffer
				input = ""
				//and clear the user input part of
				//the screen.
				for i := xMin; i < x; i++ {
					t.SetCell(i, y, ' ', Bg, Fg)
				}
				x = xMin
				err = t.Flush()
				if err != nil {
					return
				}
			case t.KeyBackspace, 0x7f:
				//If the user presses backspace,
				//then we need to remove the most
				//recent character, and decrement
				//x, but make sure not to go past
				//where we're allowed.
				if x > xMin {
					x--
					input = input[:len(input)-1]
					t.SetCell(x, y, ' ', Bg, Fg)
					err = t.Flush()
					if err != nil {
						return
					}
				}
				//If x is already at its minimum,
				//then we do nothing.

			default:
				input += string(ev.Ch)

				//If x is at the end of the screen,
				//then we need to scroll the buffer
				//to the right.
				if x == xMax {
					x--
					inputIterator := len(input) - 1
					for i := x; i >= xMin; i-- {
						//So, write the input
						//backward from the end
						//of the screen.
						t.SetCell(i, y, rune(input[inputIterator]), Fg, Bg)
						inputIterator--
					}
				}
				t.SetCell(x, y, ev.Ch, Fg, Bg)
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
