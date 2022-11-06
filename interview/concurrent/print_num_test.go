package concurrent

import "testing"

func Test_race(t *testing.T) {
	race(5)
}

func Test_collaborate(t *testing.T) {
	collaborate(5)
}
