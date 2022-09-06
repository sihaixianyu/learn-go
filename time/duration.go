package time

import (
	"fmt"
	"time"
)

func SinceDemo() {
	start := time.Now()

	time.Sleep(2 * time.Second)
	elapsed := time.Since(start).Seconds()

	fmt.Printf("%.2f\n", elapsed)
}
