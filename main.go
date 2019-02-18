package main

import (
	"battleship/battleship"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	server := battleship.Server{}
	server.Run()
}
