package main

import (
	bench "github.com/SashaCrofter/benchgolib"
	t "github.com/nsf/termbox-go"
	"log"
	"net"
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
	M.Chats[ActiveChat].NewString(SysPrefix+"Use '/connect <NickName> <ipv6>' to chat with a friend.", SysColor)

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

func listen(m *Manager) (err error) {
	var ln net.Listener
	ln, err = net.Listen("tcp", ":"+bench.Port)
	if err != nil {
		return
	}
	go func(ln net.Listener, m *Manager) {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println(conn.RemoteAddr().String(), err)
				continue
			}
			go handleConnection(m, conn)
		}
	}(ln, m)
	return
}

func handleConnection(m *Manager, conn net.Conn) {
	defer conn.Close()
	s, _, content, err := bench.ReceiveMessage(conn, m)
	if err != nil {
		return
	}
	c := m.ChatBySID(s.SID)
	if c == nil {
		//This should never be invoked, because
		//of the previous error catching block.
		return
	}
	c.NewString(InPrefix+content, InColor)
	if m.Chats[ActiveChat] == c {
		//If c is the active chat, then we can
		//update the history.
		showHistory(c.History)
	}
	//Otherwise, there is no need.
}
