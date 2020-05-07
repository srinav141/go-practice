package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	m := make(map[int]int)
	exp := len(nums) / 2

	for _, v := range nums {
		if val, ok := m[v]; ok {

			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}
	fmt.Println(m)
	for key, value := range m {
		if value > exp {
			return key
		}
	}
	return -1
}

func main() {
	r := []int{2, 2, 1, 1, 1, 2, 2}
	res := majorityElement(r)
	fmt.Println(res)
}

/*
Boyer-Moore Voting Algorithm
class Solution:
    def majorityElement(self, nums):
        count = 0
        candidate = None

        for num in nums:
            if count == 0:
                candidate = num
            count += (1 if num == candidate else -1)

        return candidate

*/
