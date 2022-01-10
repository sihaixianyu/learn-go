package channel

import "fmt"

func pipline1() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-naturals)
	}
}

func pipline2() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}

func pipline3() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)

	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}
func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
