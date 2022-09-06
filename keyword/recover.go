package keyword

import "fmt"

func safelyDo() {
	do := func() {
		panic("do something wrong!")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover from panic:", err)
		}
	}()

	do()
}
