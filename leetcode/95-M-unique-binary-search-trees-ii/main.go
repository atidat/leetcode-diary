package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generator(head, tail int) []*TreeNode {
	if head > tail {
		return []*TreeNode{nil}
	}

	repos := make([]*TreeNode, 0)
	for k := head; k <= tail; k++ {
		leftTree := generator(head, k-1)
		rightTree := generator(k+1, tail)

		for _, left := range leftTree {
			for _, right := range rightTree {
				node := TreeNode{
					Val:   k,
					Left:  left,
					Right: right,
				}
				repos = append(repos, &node)
			}
		}
	}
	return repos
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generator(1, n)
}

func main() {
}
