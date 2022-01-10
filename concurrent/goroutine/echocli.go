package goroutine

import (
	"log"
	"net"
	"os"
)

func echoReq() {
	conn, err := net.Dial("tcp", "localhost:8900")
	if err != nil {
		log.Fatal()
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
