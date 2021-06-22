package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	flag, _ := iterate(root)
	return flag
}

func iterate(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	lf, lh := iterate(node.Left)
	rf, rh := iterate(node.Right)
	if lf && rf {
		flag, maxh := check(lh, rh)
		if flag {
			return flag, maxh+1
		}
	}
	return false, -1
}

func check(l, r int) (bool, int) {
	if l >= r && l - r <= 1 {
		return true, l
	}
	if r >= l && r - l <= 1 {
		return true, r
	}
	return false, -1
}