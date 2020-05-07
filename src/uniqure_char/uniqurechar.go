package main

import (
	"fmt"
	"strings"
)

func firstUniqChar(s string) int {
	for i, v := range s {
		var res = strings.Count(s, string(v))
		if res == 1 {
			return i
		}
	}
	return -1
}

func main() {
	res := firstUniqChar("leetcode")
	fmt.Println(res)
}
