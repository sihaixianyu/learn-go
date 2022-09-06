package net

import (
	"fmt"
	"net"
	"sync"
)

func clientHello(msg string, wg *sync.WaitGroup) {
	defer wg.Done()

	client, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	_, err = client.Write([]byte(msg))
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1)
	_, err = client.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("cli-%s got resp for cli-%c, info %v\n", msg, buf[0], client.LocalAddr().String())
}

func serverHello() {
	monitor, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer monitor.Close()

	for {
		client, err := monitor.Accept()
		if err != nil {
			panic(err)
		}

		go handle(client)
	}
}

func handle(client net.Conn) {
	buf := make([]byte, 1)
	_, err := client.Read(buf)
	if err != nil {
		panic(err)
	}

	_, err = client.Write([]byte{buf[0]})
	if err != nil {
		panic(err)
	}

	client.Close()
}
