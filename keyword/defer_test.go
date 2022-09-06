package keyword

import (
	"fmt"
	"testing"
)

func TestModifyAnonymousRet(t *testing.T) {
	nums := make([]int, 0)
	nums = append(nums, 0)

	nums_1 := ModifyAnonymousRet(nums)

	fmt.Println(nums, nums_1)
}

func TestModifyNamedRet(t *testing.T) {
	nums := make([]int, 0)
	nums = append(nums, 0)

	nums_1 := ModifyNamedRet(nums)

	fmt.Println(nums, nums_1)
}

func TestMultipleDefer(t *testing.T) {
	MultipleDefer()
}

func TestNestedDefer(t *testing.T) {
	NestedDefer()
}

func TestDeferAfterPanic(t *testing.T) {
	DeferAfterPanic()
}
