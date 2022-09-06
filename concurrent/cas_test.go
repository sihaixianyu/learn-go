package concurrent

import (
	"testing"
)

func TestChangeWithoutAtomic(t *testing.T) {
	ModifyWithoutAtomic()
}

func TestChangeWithAtomic(t *testing.T) {
	ModifyWithAtomic()
}
