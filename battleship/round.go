package battleship

import (
	"log"
	"time"
)

const (
	BOARD_SIZE = 9
)

type Round struct {
	playerA, playerB *Player
	nextPlayer       *Player
	gameStarted      time.Time
}

func (r *Round) IsWaiting() bool {
	return r.playerA == nil || r.playerB == nil
}

func (r *Round) StartRound() {
	if r.IsWaiting() {
		return
	}

	r.nextPlayer = r.playerA
	r.gameStarted = time.Now()

	log.Printf("Starting round at %s", r.gameStarted.String())
	r.broadcast("Starting game round!")
}

func (r *Round) AddPlayer(p *Player) {
	switch {
	case r.playerA == nil:
		r.playerA = p
		r.playerA.AddShips()
		r.playerA.Send(NewDrawBoard(r.playerA.ships))

	case r.playerB == nil:
		r.playerB = p
		r.playerB.AddShips()
		r.playerB.Send(NewDrawBoard(r.playerB.ships))
	}
}

func (r *Round) broadcast(message interface{}) {
	if r.playerA != nil {
		l, err := r.playerA.Send(message)
		if err != nil {
			log.Print(l, err)
		}
	}
	if r.playerA != nil {
		l, err := r.playerB.Send(message)
		if err != nil {
			log.Print(l, err)
		}
	}
}

func (r *Round) End(reason string) {
	m := NewAnnouncement(reason)
	if conn := r.playerA.conn; conn != nil {
		r.playerA.Send(m)
		(*conn).Close()
	}
	if conn := r.playerB.conn; conn != nil {
		r.playerB.Send(m)
		(*conn).Close()
	}
}

func (r *Round) onError(err error) {
	r.End(err.Error())
}
