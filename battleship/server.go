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
	rounds []*Round
}

func (s *Server) Run() {
	addr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:"+strconv.Itoa(SERVER_PORT))
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
	player := Player{conn: *c}

	for _, round := range s.rounds {
		if round.IsWaiting() {
			if round.playerA.conn == nil {
				continue
			}

			connBuddy := round.playerA.conn

			log.Printf("Adding %v to an existing round with %v.", c, connBuddy)
			round.AddPlayer(&player)
			round.StartRound()
			return
		}
	}

	s.newRound(&player)
}

func (s *Server) newRound(p *Player) {
	fmt.Printf("Starting new round with player %s", p.conn)

	round := Round{}
	round.AddPlayer(p)

	s.rounds = append(s.rounds, &round)

	m := NewAnnouncement("Waiting for another player to join...")
	p.Send(m)
}

func (s *Server) endRound(r *Round) {
	rounds := make([]*Round, len(s.rounds))
	for _, round := range s.rounds {
		if &*round != &*r {
			rounds = append(rounds, round)
		}
	}
	s.rounds = rounds
}
