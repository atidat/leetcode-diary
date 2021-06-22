package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	five := ListNode{5, nil}
	four := ListNode{4, &five}
	three := ListNode{3, &four}
	two := ListNode{2, &three}
	//head := ListNode{1, &two}
	reorderList(&two)
	//reorderList(&head)
}

// 1、找到中间点
// 2、反转后半段链表的顺序
// 3、前后半段链表跳转拼接新链表

func reorderList(head *ListNode) {
	//judge head or head.Next == nil
	if head == nil || head.Next == nil {
		return
	}

	mid := getMid(head)
	fmt.Println("mid num is", mid.Val)

	var newHead *ListNode = nil
	for mid != nil {
		mid, mid.Next, newHead = mid.Next, newHead, mid
	}

	join(head, newHead)
	prt(head)
	for newHead != nil {
		head, head.Next, newHead, newHead.Next = head.Next, newHead, newHead.Next, head.Next
	}
}

func getMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp := slow.Next
	slow.Next = nil
	return tmp
}

func prt(head *ListNode) {
	for tmp := head; tmp != nil; {
		fmt.Printf("%d -> ", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println()
}
