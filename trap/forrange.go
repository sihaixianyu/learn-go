package trap

import "fmt"

func infLoop() {
	nums := []int{1, 2, 3}

	for _, v := range nums {
		nums = append(nums, v)
	}

	fmt.Println(nums)
}

func copyAddr() {
	s1 := []int{1, 2, 3}
	var s2 []*int

	for _, val := range s1 {
		s2 = append(s2, &val)
	}

	for _, val := range s2 {
		fmt.Printf("%d", *val)
	}
}

func randomTraverse() {
	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	for k, v := range hash {
		println(k, v)
	}
}
