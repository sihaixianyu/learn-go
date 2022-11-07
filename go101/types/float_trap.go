package types

import (
	"fmt"
	"math"
)

func isInfEqual() {
	a, b, c := 2.0, 1.0, 0.0
	x, y := a/c, b/c

	fmt.Println(x, y)
	// Here, x and y are equal unexpectedly
	fmt.Println(x == y)
}

func isNaNEqual() {
	x := math.NaN()
	y := math.Sqrt(-1.0)

	fmt.Println(x, y)
	// No value is euqal to NaN, not even NaN iteself
	fmt.Println(x == y)
}
