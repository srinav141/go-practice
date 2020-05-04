package main

import (
	"fmt"
	"strconv"
)

func findComplement(num int) int {

	comp := strconv.FormatInt(int64(num), 2)

	result := ""
	for _, r := range comp {
		var i, _ = strconv.ParseInt(string(r), 10, 32)
		if i == 0 {
			result += "1"
		} else {
			result += "0"
		}

	}

	res, _ := strconv.ParseInt(result, 2, 32)

	return int(res)

}

func main() {
	n := findComplement(5)
	fmt.Println(n)
}
