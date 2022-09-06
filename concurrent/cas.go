package concurrent

import (
	"fmt"
	"sync/atomic"
	"time"
)

var routineNum int
var optNum int

func init() {
	routineNum = 6
	optNum = 1000
}

func ModifyWithoutAtomic() {
	cnt := int32(0)

	for i := 0; i < routineNum; i++ {
		go func(i int) {
			for j := 0; j < optNum; j++ {
				cnt += 1
				fmt.Printf("Cnt was modified to %d by goroutine: %d\n", cnt, i)
			}
		}(i)
	}

	for cnt != 6000 {
		time.Sleep(time.Second)
	}
	fmt.Println("Done!")
}

func ModifyWithAtomic() {
	cnt := int32(0)

	for i := 0; i < routineNum; i++ {
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
		}(i)
	}

	for cnt != int32(routineNum*optNum) {
		time.Sleep(time.Second)
	}
	fmt.Println("Done!")
}
