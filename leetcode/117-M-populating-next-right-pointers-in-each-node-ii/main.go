package main

func main() {

}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	 Val int
	 Left *Node
	 Right *Node
	 Next *Node
}

/*
class Solution:
    def connect(self, root: 'Node') -> 'Node':
        if not root:
            return None
        queue = [root]
        while queue:  # 遍历每一层
            length = len(queue)
            tail = None # 每一层维护一个尾节点
            for i in range(length): # 遍历当前层
                curnode = queue.pop(0)
                if tail:
                    tail.next = curnode # 让尾节点指向当前节点
                tail = curnode # 让当前节点成为尾节点
                if curnode.left : queue.append(curnode.left)
                if curnode.right: queue.append(curnode.right)
        return root
*/
func connect(root *Node) *Node {
	cache := make([]*Node, 0)
	cache = append(cache, root)

	for len(cache) != 0 {
		var tail *Node

		cl := len(cache)
		for i := 0; i < cl; i++ {
			cur := cache[i]
			if tail != nil {
				tail.Next = cur
			}
			tail = cur
			if cur.Left != nil {
				cache = append(cache, cur.Left)
			}
			if cur.Right != nil {
				cache = append(cache, cur.Right)
			}
		}
		cache = cache[cl:]
	}

	return root
}

