package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	cur := head
	next := head.Next
	for next != nil {
		cur = next
		next = cur.Next

		
	}
}

func main() {

}
