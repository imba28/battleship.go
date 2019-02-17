package battleship

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

const (
	SERVER_PORT = 7777
)

type Server struct {
	rounds []Round
}

func (s *Server) Run() {
	addr, err := net.ResolveTCPAddr("tcp4",  "0.0.0.0:" + strconv.Itoa(SERVER_PORT))
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening on port %d", SERVER_PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go s.handleClient(&conn)
	}
}

func (s *Server) handleClient(c *net.Conn) {
	conn := *c
	player := Player{conn: c}

	for _, round := range s.rounds {
		if round.IsWaiting() {
			log.Printf("Adding %s to an existing round.", conn.RemoteAddr().String())
			round.AddPlayer(&player)
			round.StartRound()
		}
	}

	s.newRound(&player)
}

func (s *Server) newRound(p *Player) {
	conn := *p.conn
	fmt.Printf("Starting new round with player %s", conn.RemoteAddr().String())
	round := Round{playerA: p}
	s.rounds = append(s.rounds, round)
}