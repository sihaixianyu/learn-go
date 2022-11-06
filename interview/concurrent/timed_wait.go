package concurrent

import (
	"sync"
	"time"
)

func TimedWait(wg *sync.WaitGroup, t time.Duration) bool {
	signal := make(chan int)

	go func() {
		wg.Wait()
		signal <- 0
	}()

	go func() {
		time.Sleep(t)
		signal <- 1
	}()

	status := <-signal

	return status == 0
}

func TimedWait2(wg *sync.WaitGroup, t time.Duration) bool {
	signal := make(chan int)
	timer := time.NewTimer(t)

	go func() {
		wg.Wait()
		signal <- 0
	}()

	select {
	case <-signal:
		return true
	case <-timer.C:
		return false
	}
}
