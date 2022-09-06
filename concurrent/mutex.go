package concurrent

import (
	"fmt"
	"sync"
	"time"
)

func MutexExample1() {
	var mu sync.Mutex
	wg := sync.WaitGroup{}

	fmt.Println("Locked")
	mu.Lock()
	defer mu.Unlock()

	for i := 1; i < 3; i++ {

		go func(i int) {
			wg.Add(1)
			defer wg.Done()

			fmt.Println("Not lock: ", i)

			mu.Lock()
			defer mu.Unlock()

			fmt.Println("Lock: ", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock: ", i)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlocked")

	wg.Wait()
}
