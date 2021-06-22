package main

import "fmt"

func main() {
	root := build()
	fmt.Println(distributeCoins(root))

	root1 := build1()
	fmt.Println(distributeCoins(root1))

	root2 := build2()
	fmt.Println(distributeCoins(root2))
}

func build() *TreeNode {
	snd2 := TreeNode{3, nil, nil}
	fst1 := TreeNode{0, nil, &snd2}
	fst2 := TreeNode{0, nil, nil}
	root := TreeNode{1, &fst1, &fst2}
	return &root
}

func build1() *TreeNode {
	fst1 := TreeNode{0, nil, nil}
	root := TreeNode{2, &fst1, nil}
	return &root
}

func build2() *TreeNode {
	fth1 := TreeNode{2, nil, nil}
	trd1 := TreeNode{0, &fth1, nil}
	snd1 := TreeNode{0, &trd1, nil}
	fst1 := TreeNode{0, &snd1, nil}
	root := TreeNode{3, &fst1, nil}
	return &root
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type st struct {
	t *TreeNode
	rd bool // rd表示是否处理过右子树
}

func distributeCoins(root *TreeNode) int {
	var res int
	sts := make([]st, 1)
	sts[0] = st{root, false}

	node := root
	for len(sts) != 0 {
		if node == nil {
			pop(&sts)
			if len(sts) == 0 {
				break
			}
			clear(&sts, &res)

			node = top(&sts).t.Right
			continue
		}

		for node.Left != nil {
			push(&sts, st{node.Left, false})
			node = node.Left
		}

		if node.Right == nil {  // 叶节点
			clear(&sts, &res)
			if len(sts) == 0 {
				break
			}
			sts[len(sts)-1].rd = true
			node = top(&sts).t.Right
			push(&sts, st{node, false})
		} else {
			sts[len(sts)-1].rd = true
			push(&sts, st{node.Right, false})
			node = node.Right
		}
	}

	return res
}

func top(s *[]st) st {
	return (*s)[len(*s)-1]
}

func pop(s *[]st) {
	*s = (*s)[:len(*s)-1]
}

func push(s *[]st, ss st) {
	*s = append(*s, ss)
}

func setPar(s *[]st, val int) {
	pop(s)
	if len(*s) > 0 {
		(*s)[len(*s)-1].t.Val += val
	}

}

func clear(s *[]st, res *int) {
	node := (*s)[len(*s)-1]
	node.rd = true
	for node.rd || node.t.Right == nil{
		if node.t.Val >= 1 {
			*res += node.t.Val - 1
		} else {
			*res += 1 - node.t.Val
		}
		setPar(s, node.t.Val - 1)
		if len(*s) == 0 {
			return
		}
		node = top(s)
	}
}
