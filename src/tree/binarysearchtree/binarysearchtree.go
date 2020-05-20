package binarysearchtree

import "fmt"

//TreeNode node of the BST
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Create funxtion to create node
func Create(n []int) *TreeNode {

	head := TreeNode{Val: n[0]}

	for i := 1; i < len(n); i++ {
		insert(&head, n[i])
	}

	return &head

}

func insert(h *TreeNode, v int) {

	p := h

	for &p.Val != nil {

		if v < p.Val {

			if p.Left == nil {
				p.Left = &TreeNode{Val: v}
				return

			}

			p = p.Left
		}
		if v > p.Val {
			if p.Right == nil {
				p.Right = &TreeNode{Val: v}
				return
			}
			p = p.Right
		}

	}

}

//PrintInOrder traversal
func PrintInOrder(node *TreeNode) {

	if node == nil {
		return
	}

	PrintInOrder(node.Left)
	fmt.Println(node.Val)
	PrintInOrder(node.Right)

}

//PrintPreOrder traversal
func PrintPreOrder(node *TreeNode) {

	if node == nil {
		return
	}
	fmt.Println(node.Val)
	PrintInOrder(node.Left)

	PrintInOrder(node.Right)

}

//PrintPostder traversal
func PrintPostder(node *TreeNode) {

	if node == nil {
		return
	}

	PrintInOrder(node.Left)

	PrintInOrder(node.Right)
	fmt.Println(node.Val)

}

var count = 0

//KthSmallest retruns int
func KthSmallest(root *TreeNode, k int) *TreeNode {
	if root == nil {
		return nil
	}

	r := KthSmallest(root.Left, k)

	if r != nil {
		return root.Left
	}
	count++
	if k == count {
		return root
	}
	return KthSmallest(root.Right, k)

}
