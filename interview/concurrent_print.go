package interview

import (
	"fmt"
	"sync"
)

var N = 5

// 竞争型
func race() {
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

// 协同型
func collaborate() {
	chans := make([]chan bool, N+1)
	for i := 0; i <= N; i++ {
		chans[i] = make(chan bool)
	}

	for i := 1; i <= N; i++ {
		go func(id int) {
			// 当前协程阻塞，直到收到前一个协程的信号
			<-chans[id-1]
			// 打印自己对应的序号
			fmt.Println(id)
			// 通知下一个协程干活
			chans[id] <- true
		}(i)
	}
	chans[0] <- true

	// 等待最后一个协程完成工作
	<-chans[N]
}
