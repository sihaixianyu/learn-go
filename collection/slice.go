package collection

import "fmt"

func RangeAppend() {
	nums := []int{1, 2, 3}
	fmt.Printf("before range append: %v\n", nums)
	
	fmt.Println("range traverse begin...")
	// 并不会无限循环
	for _, num := range nums {
		fmt.Println(num)
		nums = append(nums, num)
	}
	fmt.Println("range traverse end...")

	fmt.Printf("after range append: %v\n", nums)
}