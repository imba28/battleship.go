package battleship

import (
	"fmt"
	"io"
)

type Player struct {
	conn  io.ReadWriteCloser
	ships []Ship
}

func (p *Player) Send(message interface{}) (int, error) {
	// todo
	if p.conn == nil {
		return 0, nil
	}

	var b []byte

	switch message.(type) {
	case string:
		m := NewAnnouncement(message.(string))
		b = []byte(m.String())
	case fmt.Stringer:
		b = []byte(message.(fmt.Stringer).String())
	}

	fmt.Println(string(b))

	return p.conn.Write(b)
}

func (p *Player) AddShips() {
	shipTypes := []rune{SHIP_CARRIER, SHIP_BATTLESHIP, SHIP_DESTROYER, SHIP_SUBMARINE}
	ships := make([]Ship, len(shipTypes))

	blockedTiles := []Coordinate{}

	for i, shipType := range shipTypes {
		ship := NewRandomShip(shipType, &blockedTiles)
		ships[i] = ship
		blockedTiles = append(blockedTiles, ship.coordinates...)
	}

	p.ships = ships
}
