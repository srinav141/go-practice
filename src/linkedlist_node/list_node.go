package main

import "fmt"

//ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {

	if head != nil {
		curr := head
		odd := head
		even := head.Next
		fe := even

		for curr != nil && curr.Next != nil {
			odd.Next = even.Next
			odd = odd.Next
			if odd != nil {
				even.Next = odd.Next
				even = even.Next
				curr = odd
			}

		}

		curr.Next = fe
	}

	return head
}

func createLinkedList(n []int) *ListNode {

	head := ListNode{Val: n[0], Next: nil}
	curr := &head

	for i := 1; i < len(n); i++ {
		node := ListNode{Val: n[i], Next: nil}
		curr.Next = &node
		curr = curr.Next

	}

	return &head
}

func main() {

	s := []int{1, 2, 3, 4}
	res := createLinkedList(s)
	r := oddEvenList(res)

	for r != nil {
		fmt.Println(r.Val)
		r = r.Next

	}
}
