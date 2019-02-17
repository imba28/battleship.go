package battleship

import (
	"log"
	"time"
)

type Round struct {
	playerA, playerB *Player
	nextPlayer *Player
	gameStarted time.Time
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
}

func (r *Round) AddPlayer(p *Player) {
	r.playerB = p
}