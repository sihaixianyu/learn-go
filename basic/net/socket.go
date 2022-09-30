package net

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func StartServer() {
	server, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
		}

		go func(conn net.Conn) {
			defer conn.Close()

			for {
				reader := bufio.NewReader(conn)
				var buf [128]byte

				n, err := reader.Read(buf[:])
				if err != nil {
					fmt.Println("server read from conn error: ", err)
					return
				}

				fmt.Println("server receive msg: ", string(buf[:n]))

				_, err = conn.Write([]byte("OK"))
				if err != nil {
					fmt.Println("server write to conn error: ", err)
					break
				}
			}
		}(conn)

		time.Sleep(5 * time.Second)
	}
}

func StartClient() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("dial error: ", err)
		return
	}

	msg := "Hello server, this is dial from client"

	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("client write to conn error: ", err)
		return
	}

	var buf [1024]byte

	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("client read from conn error: ", err)
	}

	fmt.Println("client receive msg: ", string(buf[:n]))
}
