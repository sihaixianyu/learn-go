package keyword

import "fmt"

func newDemo1() {
	numAddr := new(int)
	*numAddr = 1
	fmt.Printf("addr: %d, val: %d\n", numAddr, *numAddr)
}

func newDemo2() {
	var numAddr *int
	// Panic here, because we didn't alloc memory to pointer numAddr
	*numAddr = 1
	fmt.Printf("addr: %d, val: %d\n", numAddr, *numAddr)
}

func newDemo3() {
	var numAddr *[]int

	numAddr = new([]int)
	*numAddr = make([]int, 0)
	fmt.Printf("addr: %d, val: %d\n", numAddr, *numAddr)

	numAddr = new([]int)
	*numAddr = make([]int, 5)
	fmt.Printf("addr: %d, val: %d\n", numAddr, *numAddr)
}
