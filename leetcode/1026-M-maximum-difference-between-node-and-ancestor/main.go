package main

import (
	"fmt"
)

func main() {
	root := build1()
	fmt.Println(maxAncestorDiff(root))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func build1() *TreeNode {
	trd1 := TreeNode{4, nil, nil}
	trd2 := TreeNode{7, nil, nil}
	trd3 := TreeNode{13, nil, nil}
	snd1 := TreeNode{1, nil, nil}
	snd2 := TreeNode{6, &trd1, &trd2}
	snd3 := TreeNode{14, &trd3, nil}
	fst1 := TreeNode{3, &snd1, &snd2}
	fst2 := TreeNode{10, nil, &snd3}
	root := TreeNode{8, &fst1, &fst2}
	return &root
}



func maxAncestorDiff(root *TreeNode) int {
	res := 0
	dfs(root.Left, root.Val, &res, []int{})
	dfs(root.Right, root.Val, &res, []int{})
	return res
}

// 返回
func dfs(node *TreeNode, par int, res *int, a []int) {
	if node == nil {
		return
	}
	al := make([]int, 0)
	*res = mmax(abs(*res), abs(node.Val - par))
	al = append(al, node.Val - par)

	for _, v := range a {
		*res = mmax(abs(*res), abs(node.Val + v - par))
		al = append(al, node.Val + v - par)
	}

	fmt.Println()

	dfs(node.Left, node.Val, res, al)
	dfs(node.Right, node.Val, res, al)
}



	func mmax(a, b int) int {
		if a < b { return b }
		return a
	}

	func abs(a int) int {
		if a < 0 { return 0 - a }
		return a
	}