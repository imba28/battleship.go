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

	r.broadcast(NewDrawBoard())
}

func (r *Round) AddPlayer(p *Player) {
	switch {
	case r.playerA == nil:
		r.playerA = p
	case r.playerB == nil:
		r.playerB = p
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
