package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	sliceA, sliceB := make([]*ListNode, 0), make([]*ListNode, 0)
	for headA != nil || headB != nil {
		if headA != nil {
			sliceA = append([]*ListNode{headA}, sliceA...)
			headA = headA.Next

		}
		if headB != nil {
			sliceB = append([]*ListNode{headB}, sliceB...)
			headB = headB.Next
		}
	}

	lenSlice := len(sliceA)
	if len(sliceA) < len(sliceB) {
		lenSlice = len(sliceB)
	}

	i := 0
	for ; i < lenSlice; i++ {
		if sliceA[i] != sliceB[i] {
			break
		}
	}

	if i <= 0 {
		return nil
	}
	return sliceA[i-1]
}

func main() {

}
