package concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var routineNum int
var optNum int

func init() {
	routineNum = 10
	optNum = 100
}

func ModifyWithoutAtomic() {
	cnt := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < routineNum; i++ {
		wg.Add(1)

		go func(i int) {
			for j := 0; j < optNum; j++ {
				cnt += 1
				fmt.Printf("Cnt was modified to %d by goroutine: %d\n", cnt, i)
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Done!")
}

func ModifyWithAtomic() {
	cnt := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < routineNum; i++ {
		wg.Add(1)

		go func(i int) {
			for j := 0; j < optNum; j++ {
				for {
					ok := atomic.CompareAndSwapInt32(&cnt, cnt, cnt+1)
					if ok {
						fmt.Printf("Cnt was modified to %d by goroutine: %d\n", cnt, i)
						break
					}
				}
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Done!")
}
