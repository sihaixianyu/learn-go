package concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRWLock(t *testing.T) {
	n := 30
	state := 0
	var lock RWLock
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)

		if i%2 == 0 {
			go func(id int) {
				lock.WLock()
				defer lock.WUnlock()

				state = id
				fmt.Printf("%02d write state=%d\n", id, state)

				wg.Done()
			}(i)
			continue
		}

		go func(id int) {
			lock.RLock()
			defer lock.RUnlock()

			fmt.Printf("%02d read state=%d\n", id, state)

			wg.Done()
		}(i)

		time.Sleep(time.Millisecond)
	}

	wg.Wait()
}
