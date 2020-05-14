package main

import "fmt"

func singleNonDuplicate(nums []int) int {

	m := make(map[int]bool)
	single := 0
	for _, v := range nums {

		m[v] = !m[v]

	}

	for k, v := range m {
		if v {
			single = k
		}
	}
	return single
}

func main() {

	s := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	s = []int{3, 3, 7, 7, 10, 11, 11}
	res := singleNonDuplicate(s)
	fmt.Println(res)
}
