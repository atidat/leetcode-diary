package main

import (
	"fmt"
)

func main() {
	fmt.Println(delNodes(build(), []int{3, 5}))
	fmt.Println(delNodes(build(), []int{1, 3, 5}))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func build() *TreeNode {
	snd1 := &TreeNode{4, nil, nil}
	snd2 := &TreeNode{5, nil, nil}
	snd3 := &TreeNode{6, nil, nil}
	snd4 := &TreeNode{7, nil, nil}

	fst1 := &TreeNode{2, snd1, snd2}
	fst2 := &TreeNode{3, snd3, snd4}

	root := &TreeNode{1, fst1, fst2}
	return root
}
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	dels := make(map[int]bool, len(to_delete))
	for _, v := range to_delete {
		dels[v] = true
	}

	res := make([]*TreeNode, 0)
	dfs(root, dels, &res)
	r := post(&res)
	return r
}

func post(trees *[]*TreeNode) []*TreeNode {
	if len(*trees) < 2 {
		return *trees
	}
	res := make([]*TreeNode, 0)
	l, r := 0, 1
	for ; r < len(*trees); {
		if (*trees)[l].Left != nil &&(*trees)[l].Left.Val == (*trees)[r].Val ||
			(*trees)[l].Right != nil && (*trees)[l].Right.Val == (*trees)[r].Val {
			//
		} else {
			res = append(res, (*trees)[r])
		}
		l++
		r++
	}
	res = append(res, (*trees)[0])
	return res
}

func dfs(node *TreeNode, dels map[int]bool, res *[]*TreeNode) *TreeNode {
	if node == nil {
		return node
	}

	var isin bool
	if _, ok := dels[node.Val]; ok {
		isin= true
	}
	if !isin {
		push(res, node)
	}

	if node.Left != nil {
		if dfs(node.Left, dels, res) == nil {
			node.Left = nil
		}
	}

	if node.Right != nil {
		if dfs(node.Right, dels, res) == nil {
			node.Right = nil
		}
	}

	if isin {
		return nil
	}
	return node
}

func push(res *[]*TreeNode, t *TreeNode) {
	*res = append(*res, t)
}