package keyword

import (
	"fmt"
	"testing"
)

var s Student

func init() {
	s = Student{
		name:  "ztx",
		class: 1,
		age:   23,
	}
}

func TestPhone_Upgrade(t *testing.T) {
	fmt.Println(s)

	s.Upgrade()
	fmt.Println(s)

	(&s).Upgrade()
	fmt.Println(s)
}

func TestPhone_Grow(t *testing.T) {
	fmt.Println(s)

	s.Grow()
	fmt.Println(s)

	(&s).Grow()
	fmt.Println(s)
}
