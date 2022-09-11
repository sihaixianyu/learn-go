package concurrent

import (
	"fmt"
	"time"
)

// In this example, the worker has been synchronized through channel.
func chanForSync() {
	done := make(chan int)

	for i := 0; i < 2; i++ {
		go func(i int) {
			fmt.Printf("pre-work[%d] start...\n", i)
			if i == 0 {
				time.Sleep(time.Second)
			}
			done <- i
			fmt.Printf("pre-work[%d] done!\n", i)
		}(i)

	}

	fmt.Println("work start...")

	fmt.Printf("pre-work[%d] is over\n", <-done)
	fmt.Printf("pre-work[%d] is over\n", <-done)

	fmt.Println("work done!")
}
