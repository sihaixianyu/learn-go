package concurrent

import (
	"log"
	"net"
	"os"
)

func Call2() {
	conn, err := net.Dial("tcp", "localhost:8900")
	if err != nil {
		log.Fatal()
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
