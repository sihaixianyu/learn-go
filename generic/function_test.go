package generic

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	s1 := &Student{
		Age: 11,
	}

	s2 := &Student{
		Age: 12,
	}

	fmt.Println(Max(s1, s2))
}
