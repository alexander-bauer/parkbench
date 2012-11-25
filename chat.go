package main

import (
	"crypto/rand"
	"crypto/rsa"
	bench "github.com/SashaCrofter/benchgolib"
	t "github.com/nsf/termbox-go"
)

type Manager struct {
	key     *rsa.PrivateKey  //TEMP; in-memory private key
	ConfDir string           //The ParkBench directory
	Chats   map[string]*Chat //All open chats, mapped by partner
}

func NewManager(confDir string) (m *Manager) {
	return &Manager{
		ConfDir: confDir,
		Chats:   make(map[string]*Chat, 0),
	}
}

//No-op, intended to be replaced by Chat.Connect.
func (m *Manager) AddSession(s *bench.Session) error {
	return nil
}

//Goes through all existing chats until the SID matches.
func (m *Manager) SessionByID(sid uint64) *bench.Session {
	for _, v := range m.Chats {
		if v.Session.SID == sid {
			return v.Session
		}
	}
	return nil
}

func (m *Manager) PrivateKey() *rsa.PrivateKey {
	if m.key == nil {
		var err error
		m.key, err = rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			panic(err)
		}
		//This doesn't catch any errors, possibly
		//resulting in runtime errors.
	}
	return m.key
}

type Chat struct {
	manager *Manager       //The Manager that this chat belongs to
	Partner string         //Identifier for the chat partner
	History [][]t.Cell     //The list of all items in the chat window
	Session *bench.Session //The current Session for the chat
}

//Returns a *Chat with the Partner and History objects initialized. The Session field is not initialized, and must be made using Chat.Connect().
func (m *Manager) NewChat(partner string) *Chat {
	return &Chat{
		manager: m,
		Partner: partner,
		History: make([][]t.Cell, 0),
	}
}

//Sets the Session field of the given Chat to the returned Session, returning any errors. c.Session may be nil afterward.
func (c *Chat) Connect(remote string) (err error) {
	c.Session, err = bench.NewSession(remote, c.manager)
	return
}
