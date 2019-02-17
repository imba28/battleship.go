package battleship

import "net"

type Player struct {
	conn *net.Conn
	ships *[]Ship
}
