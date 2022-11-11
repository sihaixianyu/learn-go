package main

import "fmt"

func testAssign(nums *[]int) {
	*nums = []int{1, 2, 3}
}

func main() {
	nums := []int{}
	fmt.Println(nums)

	testAssign(&nums)
	fmt.Println(nums)
}