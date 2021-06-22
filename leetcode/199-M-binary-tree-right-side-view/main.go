package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	rights := make([]int, 0)
	if root == nil {
		return rights
	}

	eachLevelNodes := []*TreeNode{root}
	tmp := make([]*TreeNode, 0)
	for len(eachLevelNodes) != 0 {
		curNode := eachLevelNodes[0]
		if curNode.Left != nil {
			tmp = append(tmp, curNode.Left)
		}
		if curNode.Right != nil {
			tmp = append(tmp, curNode.Right)
		}

		if len(eachLevelNodes) == 1 {
			rights = append(rights, curNode.Val)
			eachLevelNodes = tmp
			tmp = []*TreeNode{}
		} else {
			eachLevelNodes = eachLevelNodes[1:]
		}
	}
	return rights
}

func main() {
}
