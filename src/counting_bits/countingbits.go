package main

import (
	"fmt"
	"math/bits"
)

func countBits(num int) []int {

	res := make([]int,0)
	for i:=0;i<=num;i++{
		n := uint(i)

		res = append(res,bits.OnesCount(n))
	}

	return res
}

func main()  {
	r := countBits(5)
	fmt.Println(r)
}