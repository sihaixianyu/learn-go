package interview

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimedWait(t *testing.T) {
	n := 10
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Printf("G %2d is working!\n", id)
		}(i)
	}

	if TimedWait(&wg, 3*time.Second) {
		fmt.Println("All the work has been done!")
		return
	}

	fmt.Println("Timeout!")
}

func TestTimedWait2(t *testing.T) {
	n := 10
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Printf("G %2d is working!\n", id)
			time.Sleep(2 * time.Second)
		}(i)
	}

	if TimedWait2(&wg, 3*time.Second) {
		fmt.Println("All the work has been done!")
		return
	}

	fmt.Println("Timeout!")
}
