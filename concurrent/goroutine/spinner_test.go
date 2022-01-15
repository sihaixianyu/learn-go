package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestFib(t *testing.T) {
	go Spinner(100 * time.Millisecond)
	const n = 45
	fibN := Fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
