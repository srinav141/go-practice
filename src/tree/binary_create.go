package main

import (
	"fmt"

	"github.com/srinav141/go-practice/src/tree/binarysearchtree"
)

func main() {
	input := []int{8, 5, 1, 7, 10, 12}
	//input = []int{1}
	res := binarysearchtree.Create(input)
	//fmt.Println(res.Left.Right.Val)
//	binarysearchtree.PrintInOrder(res)
	binarysearchtree.PrintPreOrder(res)
	//r := binarysearchtree.KthSmallest(res, 1)
	fmt.Println(res.Val)
}
