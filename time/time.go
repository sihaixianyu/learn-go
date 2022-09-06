package time

import (
	"fmt"
	"time"
)

func FormatDemo() {
	now := time.Now()
	fmt.Println(now.Format(time.UnixDate))
}
