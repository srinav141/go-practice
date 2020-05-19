package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(s, p string) []int {
	r := make([]int, 0)
	for i := 0; i < len(s); i++ {
		start := i
		end := i + len(p)
		if end <= len(s) {
			expected := s[start:end]
			if isAnagramIndex(p, expected) {

				r = append(r, start)
			}
		}
	}
	return r
}

func isAnagramIndex(actual, expected string) bool {

	if len(actual) != len(expected) {
		return false
	}

	s := strings.Split(actual, "")
	e := strings.Split(expected, "")

	sort.Strings(s)
	sort.Strings(e)
	if strings.Join(s, "") == strings.Join(e, "") {
		return true

	}

	return false
}

func main() {
	s1 := "abab"
	p := "ab"
	res := findAnagrams(s1, p)
	fmt.Println(res)
}
