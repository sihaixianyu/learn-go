package concurrent

import (
	"log"
	"net"
	"testing"
)

func Test_echo(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		//go handleConn2(conn)
		handleConn2(conn)
	}
}
