/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	queue := &NodeQueue{}
	nextQueue := &NodeQueue{}

	if root != nil {
		queue.Push(root)
	}

	ans := [][]int{}

	for !queue.Empty() {
		current := []int{}
		for !queue.Empty() {
			node := queue.Pop()
			current = append(current, node.Val)
			for _, child := range node.Children {
				nextQueue.Push(child)
			}
		}
		ans = append(ans, current)
		queue = nextQueue
		nextQueue = &NodeQueue{}
	}

	return ans
}

type NodeQueue []*Node

func (q *NodeQueue) Push(node *Node) {
	*q = append(*q, node)
}

func (q *NodeQueue) Pop() *Node {
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

func (q *NodeQueue) Empty() bool {
	return len(*q) == 0
}

// @lc code=end

