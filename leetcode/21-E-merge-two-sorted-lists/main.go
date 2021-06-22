package main

import "fmt"

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		if l1 == nil {
			return l2
		}
		return l1
	}


	var head *ListNode
	run := head
	r1, r2 := l1, l2
	for r1 != nil && r2 != nil {
		if r1.Val <= r2.Val {
			run.Next = r1
			r1 = r1.Next
		} else {
			run.Next = r2
			r2 = r2.Next
		}

	}
	if r1 == nil {
		run.Next = r2
	} else {
		run.Next = r1
	}
	doPrint(head)
	return head
}

func doPrint(h *ListNode) {
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}
