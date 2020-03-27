package main

import (
	"flag"
	"github.com/imba28/battleship/pkg/battleship"
	"log"
	"math/rand"
	"regexp"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	isClient := flag.Bool("client", true, "start as client")
	addr := flag.String("server", "127.0.0.1", "ip address of the server")
	flag.Parse()

	if *isClient {
		r := regexp.MustCompile("(?:[0-9]{0,3}\\.){3}[0-9]{1,3}$")

		if !r.MatchString(*addr) {
			log.Fatal(addr, " is not a valid TCP4 Address!")
		}

		c := battleship.Connect(*addr)
		c.Play()
	} else {
		server := battleship.Server{}
		server.Run()
	}

}
