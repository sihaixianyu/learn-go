package concurrent

import (
	"fmt"
	"sync"
)

func MutexExample1() {
	var mu sync.Mutex
	wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			mu.Lock()
			fmt.Printf("%2d got lock \n", i)

			mu.Unlock()
			fmt.Printf("%2d release lock \n", i)
		}(i)
	}

	wg.Wait()
}
