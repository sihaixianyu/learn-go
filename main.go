package main

import (
	"fmt"
	"sync"
)

var N int = 5

func main() {
	var mu sync.Mutex
	seq := 1
	wg := sync.WaitGroup{}

	for i := 1; i <= N; i++ {
		wg.Add(1)

		go func(id int) {
			for {
				mu.Lock()
				if id == seq {
					fmt.Println(seq)
					seq += 1
					break
				}
				mu.Unlock()
			}

			mu.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
}
