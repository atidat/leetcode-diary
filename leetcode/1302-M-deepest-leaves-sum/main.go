package main

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 获取最深叶子节点的和，可能有多个叶子节点处于同一深度
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	cache := make(map[int][]int)
	maxDepth := recur(root, 1, cache)
	var sum int
	for _, v := range cache[maxDepth] {
		sum += v
	}
	return sum
}

func recur(node *TreeNode, depth int, ca map[int][]int) int {
	if node == nil {
		return depth
	}

	if node.Left == nil && node.Right == nil {
		if v, ok := ca[depth]; ok {
			ca[depth] = append(v, node.Val)
		} else {
			ca[depth] = []int{node.Val}
		}
		return depth
	}

	leftDepth := recur(node.Left, depth+1, ca)
	rightDepth := recur(node.Right, depth+1, ca)
	return mmax(leftDepth, rightDepth)
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}