package battleship

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func Connect(server string) Client {
	addr := server + ":" + strconv.Itoa(SERVER_PORT)
	fmt.Printf("Connecting to %s\n", addr)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	return Client{conn: conn}
}

type Client struct {
	conn *net.TCPConn
}

func (c *Client) Play() {
	for {
		if c.conn == nil {
			log.Fatal("Disconnected!")
		}

		b := make([]byte, 128)
		_, err := c.conn.Read(b)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
		var m Message
		if err := json.Unmarshal(bytes.Trim(b, "\x00"), &m); err != nil {
			log.Println(err)
		} else {
			switch m.Name {
			case MESSAGE_INFO:
				fmt.Println(m.Body)

			case MESSAGE_DRAW_BOARD:
				c.drawBoard(c.generateBoard(m.Body.(string)))
			}
		}
	}
}

func (c *Client) generateBoard(s string) [][]string {
	board := make([][]string, BOARD_SIZE)

	for y := 0; y < BOARD_SIZE; y++ {
		row := make([]string, BOARD_SIZE)

		for x := 0; x < BOARD_SIZE; x++ {
			row[x] = " "
		}
		board[y] = row
	}

	shipType := ""
	for i := 0; i < len(s); i += 2 {
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			x, _ := strconv.Atoi(string(s[i]))
			y, _ := strconv.Atoi(string(s[i+1]))
			board[y][x] = shipType
		} else {
			shipType = string(s[i])
			i--
		}

	}

	return board
}

func (c *Client) drawBoard(b [][]string) {
	fmt.Print("    ")
	for i := 0; i < BOARD_SIZE; i++ {
		fmt.Print(string(65+i), "   ")
	}
	fmt.Print("    \n")

	for y := 0; y < BOARD_SIZE; y++ {
		fmt.Print("  ", strings.Repeat("|---", BOARD_SIZE), "|\n")
		fmt.Print(y, " ")
		for x := 0; x < BOARD_SIZE; x++ {
			fmt.Print("| ", b[y][x], " ")
		}
		fmt.Print("|\n")
	}
	fmt.Print("  ", strings.Repeat("|---", BOARD_SIZE), "|\n")
}
