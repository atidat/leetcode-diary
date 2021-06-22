package main

import (
	"math"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stk := &stack{
		buffer: []*TreeNode{root},
		l: 1,
	}

	visNodes := make(map[*TreeNode]bool)
	var result = math.MinInt32

	for !stk.empty() {
		s := stk.top()
		if s.Left != nil {
			if _, ok := visNodes[s.Left]; !ok { // s.Left未访问
				stk.push(s.Left)
				visNodes[s.Left] = true
				continue
			}
		}

		if s.Right != nil {
			if _, ok := visNodes[s.Right]; !ok { // s.Right未访问
				stk.push(s.Right)
				visNodes[s.Right] = true
				continue
			}
		}

		slv := math.MinInt32
		if s.Left != nil {
			slv = s.Left.Val
		}
		srv := math.MinInt32
		if s.Right != nil {
			srv = s.Right.Val
		}
		result = mmax(result, specMax(s.Val, slv, srv))
		if math.MinInt32 != mmax(slv, srv) {
			s.Val = mmax(s.Val, s.Val + mmax(slv, srv))
		}
		stk.pop()

	}
	// fmt.Printf("result is %d", result)
	return result
}

func specMax(ro, l, ri int) int {
	if l > 0 && ri > 0 {
		return ro + l + ri
	}
	if ro > 0 {
		return ro + mmax(0, mmax(l, ri))
	}
	return mmax(ro, mmax(l, ri))
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}

type stack struct {
	buffer []*TreeNode
	l int
}

func (s *stack)push(t *TreeNode) {
	s.buffer = append(s.buffer, t)
	s.l += 1
}

func (s *stack)pop() {
	s.l -= 1
	if s.l <= 0 {
		s.buffer = []*TreeNode{}
	} else {
		s.buffer = s.buffer[0:s.l]
	}
}

func (s *stack)top() *TreeNode {
	return s.buffer[s.l-1]
}

func (s *stack)empty() bool {
	return s.l == 0
}

