package main

import (
	"fmt"
	"math"
)

func isPerfectSquare(num int) bool {
	res := int(math.Sqrt(float64(num)))
	fmt.Println(res)

	if res*res == num {
		return true
	}
	return false
}

func main() {
	r := isPerfectSquare(16)
	fmt.Println(r)

}
