package main

import "battleship/battleship"

func main() {
	client := battleship.Connect("127.0.0.1")
	client.Play()
}
