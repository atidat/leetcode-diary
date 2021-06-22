package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	//prtTree(root, 0)

	repo := make([]int, 0)
	if root == nil {
		return repo
	}
	visited := make(map[*TreeNode]bool)
	nodes := make([]*TreeNode, 1)
	nodes[0] = root

	for len(nodes) != 0 {
		node := nodes[len(nodes)-1]
		for node.Left != nil {
			_, ok := visited[node.Left]
			if ok {
				break
			}
			nodes = append(nodes, node.Left)
			visited[node.Left] = true
			node = node.Left
		}

		//判断此时该节点是否有右子树
		//如果有右子树，则添加该该节点，并追溯其左右子树
		//如果无右子树，则剔除该节点，回溯其父节点的右子树，
		for node.Right == nil {
			repo = append(repo, node.Val)
			nodes = nodes[0 : len(nodes)-1]
			if len(nodes) == 0 {
				return repo
			}
			node = nodes[len(nodes)-1]
		}

		_, ok := visited[node.Right]
		if !ok {
			nodes = append(nodes, node.Right)
			visited[node.Right] = true
		} else {
			repo = append(repo, node.Val)
			nodes = nodes[0 : len(nodes)-1]
		}

	}
	return repo
}

func main() {
	// snd3 := TreeNode{3, nil, nil}
	// fst2 := TreeNode{2, &snd3, nil}
	// root := TreeNode{1, nil, &fst2}

	// fst2 := TreeNode{2, nil, nil}
	// fst1 := TreeNode{1, nil, nil}
	// root := TreeNode{3, &fst1, &fst2}

	snd2 := TreeNode{2, nil, nil}
	fst1 := TreeNode{1, nil, &snd2}
	root := TreeNode{3, &fst1, nil}
	fmt.Println(postorderTraversal(&root))
}

func prtTree(root *TreeNode, level int) {
	if root == nil {
		return
	}
	fmt.Printf("%d level: val is %d\n", level, root.Val)
	prtTree(root.Left, level+1)
	prtTree(root.Right, level+1)
}
