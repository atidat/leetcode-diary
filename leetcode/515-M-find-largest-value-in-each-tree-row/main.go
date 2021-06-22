package main

import (
	"fmt"
	"sort"
)

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 找到每一行的最大值
func largestValues(root *TreeNode) []int {
	dic := make(map[int][]int)
	dep := dfs(root, 0, dic)
	return mysort(dic, dep)
}

func dfs(node *TreeNode, cur int, dic map[int][]int) int {
	if node == nil {
		return cur
	}
	if v, ok := dic[cur]; !ok {
		dic[cur] = make([]int, 0)
		dic[cur] = append(dic[cur], node.Val)
	} else {
		dic[cur] = append(v, node.Val)
	}
	leftDep := dfs(node.Left, cur+1, dic)
	rightDep := dfs(node.Right, cur+1, dic)
	return mmax(leftDep, rightDep)
}

func mysort(dic map[int][]int, dep int) []int {
	res := make([]int, 0)
	for i := 0; i < dep; i++ {
		sort.Ints(dic[i])
		res = append(res, dic[i][len(dic[i])-1])
	}
	fmt.Println(res)
	return res
}

func mmax(a, b int) int {
	if a < b {
		return b
	}
	return a
}