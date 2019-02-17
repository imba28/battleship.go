package battleship

import "time"

const (
	MESSAGE_INFO       = "info"
	MESSAGE_DRAW_BOARD = "drawBoard"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func newMessage(name, body string) Message {
	return Message{Name: name, Body: body, Time: time.Now().Unix()}
}

func NewAnnouncement(body string) Message {
	return newMessage(MESSAGE_INFO, body)
}

func NewDrawBoard() Message {
	return newMessage(MESSAGE_DRAW_BOARD, "todo")
}
