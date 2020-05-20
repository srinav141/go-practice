package main

import (
	"fmt"

	"github.com/srinav141/go-practice/src/tree/binarysearchtree"
)

func main() {
	input := []int{5, 3, 6, 2, 4, 1}
	input = []int{1}
	res := binarysearchtree.Create(input)
	//fmt.Println(res.Left.Right.Val)
	binarysearchtree.PrintInOrder(res)
	r := binarysearchtree.KthSmallest(res, 1)
	fmt.Println(r.Val)
}
