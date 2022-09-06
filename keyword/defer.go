package keyword

import (
	"fmt"
)

func ModifyAnonymousRet(nums []int) []int {
	nums = append(nums, 1)

	defer func() {
		nums = append(nums, 2)
	}()

	// 匿名返回值：操作为 => ~r <- nums，此时defer函数中对nums的修改不会作用到已经被赋值的匿名返回值~r中
	return nums
}

func ModifyNamedRet(nums []int) (res []int) {
	nums = append(nums, 1)

	defer func() {
		res = append(nums, 2)
	}()

	// 命名返回值：操作为 => res <- nums，此时defer函数中对res的修改会生效
	return nums
}

func MultipleDefer() {
	printVal := func(num int) {
		fmt.Println("in printVal", num)
	}

	printPtr := func(num *int) {
		fmt.Println("in printPtr", *num)
	}

	a := 0
	defer printVal(a)
	defer printPtr(&a)

	a += 1
	fmt.Println("in invoker:", a)
}

func NestedDefer() {
	defer func() {
		defer func() {
			fmt.Println("This is inner defer")
		} ()
		fmt.Println("This is outer defer")
	} ()
	fmt.Println("This is NestedDefer")
}

func DeferAfterPanic() {
	defer func() {
		fmt.Println("defer before panic execute!")
	}()
	panic("panic in DeferAfterPanic")
}
