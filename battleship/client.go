package battleship

import (
	"fmt"
	"log"
	"net"
	"strconv"
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
		b := make([]byte, 64)
		_, err := c.conn.Read(b)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
	}
}