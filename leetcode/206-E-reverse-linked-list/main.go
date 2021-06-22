package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre, post := head, head.Next
	head.Next = nil
	for post != nil {
		tmp := post.Next

		post.Next = pre
		pre = post
		post = tmp
	}
	return pre
}

func main() {
	four := &ListNode{4, nil}
	three := &ListNode{3, four}
	two := &ListNode{2, three}
	one := &ListNode{1, two}
	head := &ListNode{0, one}
	tmp := reverseList(head)

	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}
