package main

import (
	bench "github.com/SashaCrofter/benchgolib"
	t "github.com/nsf/termbox-go"
)

type Chat struct {
	Partner string         //Identifier for the chat partner
	History [][]t.Cell //The list of all items in the chat window
	S       *bench.Session //The current Session for the chat
}
