package net

import (
	"strconv"
	"sync"
	"testing"
)

func Test_clientHello(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		msg := strconv.Itoa(i)
		go clientHello(msg, &wg)
	}

	wg.Wait()
}

func Test_serverHello(t *testing.T) {
	serverHello()
}
