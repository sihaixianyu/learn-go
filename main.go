package main

import (
	"fmt"
	"learn-go/concurrent"
	"time"
)

func main() {
	go concurrent.Spinner(100 * time.Millisecond)
	const n = 45
	fibN := concurrent.Fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
