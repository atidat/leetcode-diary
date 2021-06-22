package main

import "fmt"

func main() {
	root := buildTree1()
	fmt.Println(pseudoPalindromicPaths(root))
}

func buildTree1() *TreeNode {
	snd1 := TreeNode{3, nil, nil}
	snd2 := TreeNode{1, nil, nil}
	snd4 := TreeNode{1, nil, nil}
	fst1 := TreeNode{3, &snd1, &snd2}
	fst2 := TreeNode{1, nil, &snd4}
	root := TreeNode{2, &fst1, &fst2}
	return &root
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type st struct {
	t *TreeNode
	rok bool // ok is true, means left node has handled
}

func pseudoPalindromicPaths (root *TreeNode) int {
	// map[int]int: 节点的值 =》 该值出现的次数
	m := make(map[int]int, 1)
	r := make([]st, 1)
	var res int
	r[0] = st{root, false}
	m[root.Val] = 1

	node := root
	for len(r) != 0 {
		for node.Left != nil {
			op("+", m, node.Left.Val)
			push(&r, st{node.Left, false})
			node = node.Left
		}


		if node.Right == nil {
			if isPa(m) {
				res++
			}
			clear(m, &r)
			if len(r) == 0 {
				break
			}
			r[len(r)-1].rok = true
			node = top(&r).t.Right
			push(&r, st{node, false})
			op("+", m, node.Val)
		} else {
			r[len(r)-1].rok = true
			push(&r, st{node.Right, false})
			op("+", m, node.Right.Val)
			node = node.Right
		}
	}

	return res
}

func isPa(m map[int]int) bool {
	var sin int
	for _, v := range m {
		if v % 2 != 0 {
			sin++
		}
	}
	return sin == 1 || sin == 0
}


func op( sym string, m map[int]int, k int) {
	if v, ok := m[k]; !ok {
		if sym == "+" {
			m[k] = 1
		}
	} else {
		if sym == "+" {
			m[k] = v + 1
		} else if v > 0 {
			m[k] = v - 1
		}
	}
}

func top(r *[]st) st {
	return (*r)[len(*r)-1]
}

func pop(r *[]st) {
	*r = (*r)[:len(*r)-1]
}

func push(r *[]st, a st) {
	*r = append(*r, a)
}

func clear(m map[int]int, r *[]st) {
	tmp := top(r)
	for tmp.rok || tmp.t.Right == nil {
		op("-", m, tmp.t.Val)
		pop(r)
		if len(*r) == 0 {
			break
		}
		tmp = top(r)
	}
}