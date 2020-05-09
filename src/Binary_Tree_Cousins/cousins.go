package main

import "fmt"

/*
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {

	px, dx := getDepthAndParent(root, x)

	py, dy := getDepthAndParent(root, y)

	if px != py && dx == dy {
		return true
	}

	return false
}

func getDepthAndParent(root *TreeNode, x int) (p, d int) {

	l := root.Left
	r := root.Right

	if x%2 == 0 {
		p, d = getDepth(r, x)
	} else {
		p, d = getDepth(l, x)
	}
	return
}

func getDepth(r *TreeNode, x int) (p, d int) {
	d = 1
	v := r.Val
	parent := r
	for v != x {

		if parent.Right != nil {
			v = parent.Right.Val
			parent = parent.Right
		} else if parent.Left != nil {
			v = parent.Left.Val
			parent = parent.Left
		}
		d++

	}
	return parent.Val, d
}

func createBinarytree(s1 []int) TreeNode {

	root := TreeNode{Val: s1[0]}
	temp := &root
	r := TreeNode{Val: s1[1]}
	l := TreeNode{Val: s1[2]}
	temp.Right = &r
	fmt.Println(root.Right.Val)

	temp.Left = &l

	for i := range s1 {
		if i == 0 {

			continue
		}

		if i%2 == 0 {
			temp = temp.Left
			if i+3 < len(s1) {
				r = TreeNode{Val: s1[i+3]}
			} else {

				r = TreeNode{}
			}
			if i+4 < len(s1) {
				l = TreeNode{Val: s1[i+4]}
			} else {
				l = TreeNode{}
			}
			temp.Right = &r
			temp.Left = &l
		} else {
			temp = temp.Right

			if i+2 < len(s1) {
				r = TreeNode{Val: s1[i+2]}
			} else {
				r = TreeNode{}
			}
			if i+3 < len(s1) {
				l = TreeNode{Val: s1[i+3]}
			} else {

				l = TreeNode{}
			}
			temp.Right = &r
			fmt.Println(root.Right.Right.Val)
			temp.Left = &l
		}

	}
	return root

}

func main() {

	s1 := []int{1, 4, 5, 6}
	r := createBinarytree(s1)
	fmt.Println(r)

}
