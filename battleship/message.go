package battleship

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

const (
	MESSAGE_INFO       = "info"
	MESSAGE_DRAW_BOARD = "drawBoard"
)

type Message struct {
	Name string
	Body interface{}
	Time int64
}

func newMessage(name string, body interface{}) Message {
	return Message{Name: name, Body: body, Time: time.Now().Unix()}
}

func NewAnnouncement(body string) Message {
	return newMessage(MESSAGE_INFO, body)
}

func NewDrawBoard(ships []Ship) Message {
	b := []string{}

	for _, s := range ships {
		b = append(b, string(s.shipType))
		for _, c := range s.coordinates {
			b = append(b, strconv.Itoa(c.x))
			b = append(b, strconv.Itoa(c.y))
		}
	}
	return newMessage(MESSAGE_DRAW_BOARD, strings.Join(b, ""))
}

func (m Message) String() string {
	s, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(s)
}
