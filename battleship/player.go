package battleship

import (
	"fmt"
	"net"
)

type Player struct {
	conn  *net.Conn
	ships *[]Ship
}

func (p *Player) Send(message interface{}) (int, error) {
	conn := *p.conn
	var b []byte

	switch message.(type) {
	case string:
		m := NewAnnouncement(message.(string))
		b = []byte(m.String())
	case fmt.Stringer:
		b = []byte(message.(fmt.Stringer).String())
	}

	return conn.Write(b)
}
