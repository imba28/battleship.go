package main

import (
	"battleship/battleship"
	"log"
	"os"
	"regexp"
)

func main() {
	addr := "127.0.0.1"
	if len(os.Args) > 1 {
		addr = os.Args[1]
		r := regexp.MustCompile("(?:[0-9]{0,3}\\.){3}[0-9]{1,3}$")

		if !r.MatchString(addr) {
			log.Fatal(addr, " is not a valid TCP4 Address!")
		}
	}

	battleship.Connect(addr).Play()
}
