package battleship

import (
	"encoding/json"
	"net"
)

type Player struct {
	conn  *net.Conn
	ships *[]Ship
}

func (p *Player) Send(message Message) (int, error) {
	conn := *p.conn
	b, err := json.Marshal(message)
	if err != nil {
		return 0, err
	}

	return conn.Write(b)
}
