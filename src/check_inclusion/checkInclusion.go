package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}
	countP := [256]int{}
	countTW := [256]int{}
	for i := 0; i < len(s1); i++ {
		countP[int(s1[i])]++
		countTW[int(s2[i])]++
	}

	for i := len(s1); i < len(s2); i++ {

		if compare(countP, countTW) {
			return true

		}

		countTW[int(s2[i])]++
		countTW[int(s2[i-len(s1)])]--

	}
	if compare(countP, countTW) {
		return true

	}

	return false
}

func compare(a1, a2 [256]int) bool {

	for i := 0; i < 256; i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func checkInclusion2(s1 string, s2 string) bool {
	m := make(map[string]int)

	for _, v := range s2 {

		st := string(v)

		if val, ok := m[st]; ok {

			m[st] = val + 1

		} else {
			m[st] = 1
		}

	}

	m2 := make(map[string]int)
	for _, v := range s1 {
		st := string(v)

		if val, ok := m2[st]; ok {

			m2[st] = val + 1

		} else {
			m2[st] = 1
		}

	}

	for i := range m2 {

		if m2[i] != m[i] {
			return false
		}

	}
	return true
}

func main() {

	res := checkInclusion("a", "ab")
	fmt.Println(res)
}
