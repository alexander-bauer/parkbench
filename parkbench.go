package main

import (
	//"github.com/SashaCrofter/benchgolib"
	t "github.com/nsf/termbox-go"
	"log"
	"os/user"
	"path"
)

var (
	M *Manager
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
	usr, err := user.Current()
	if err != nil {
		return
	}
	confDir := path.Join(usr.HomeDir, ".parkbench")

	M = NewManager(confDir) //Now initialize the Manager,
	ActiveChat = "main"     //switch the chat to "main",
	M.NewChat(ActiveChat)   //and start a chat of that name.

	err = t.Init()
	if err != nil {
		return
	}
	defer t.Close()
	t.Clear(Fg, Bg)

	setString(0, 0, "ParkBench", Fg|t.AttrBold, Bg)
	M.Chats[ActiveChat].NewString(SysPrefix+"Use '/connect <ipv6>' to chat with a friend.", SysColor)

	//This will flush to the screen, as well.
	err = showHistory(M.Chats[ActiveChat].History)

	//Queue is declared as a global variable
	//in cli.go.
	go func(queue chan<- t.Event) {
		for {
			queue <- t.PollEvent()
		}
	}(Queue)

	//Now, just take user input until the user exits.
	err = loopIn(">> ", Queue)
	return
}
