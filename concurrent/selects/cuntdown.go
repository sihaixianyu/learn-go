package selects

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Racket is launching!")
}

func CountDown() {
	fmt.Println("Commencing countdown.")

	tick := time.Tick(1 * time.Second)

	for cnt := 10; cnt > 0; cnt-- {
		fmt.Println(cnt)
		<-tick
	}

	launch()
}

func CountDown2() {
	abort := make(chan struct{})
	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")

	select {
	case <-time.After(10 * time.Second):
	// Do nothing
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}

	launch()
}

func CountDown3() {
	abort := make(chan struct{})
	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for cnt := 10; cnt > 0; cnt-- {
		fmt.Println(cnt)
		select {
		case <-ticker.C:
		// Do nothing
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}

	launch()
}
