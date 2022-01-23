/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	// 方法一：BFS
	currentQ, nextQ := new(nodeQ), new(nodeQ)
	if root != nil {
		currentQ.Push(root)
	}

	for !currentQ.Empty() {
		for !currentQ.Empty() {
			node := currentQ.Pop()
			if !currentQ.Empty() {
				node.Next = currentQ.Top()
			}
			if node.Left != nil {
				nextQ.Push(node.Left)
				nextQ.Push(node.Right)
			}
		}
		currentQ = nextQ
		nextQ = new(nodeQ)
	}

	return root

	// 方法二：DFS
	// if root != nil && root.Left != nil {
	// 	connect2(root.Left, root.Right)
	// }

	// return root
}

// BFS
type nodeQ []*Node

func (q *nodeQ) Push(node *Node) {
	*q = append(*q, node)
}

func (q *nodeQ) Pop() *Node {
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

func (q *nodeQ) Top() *Node {
	return (*q)[0]
}

func (q *nodeQ) Empty() bool {
	return len(*q) == 0
}

// DFS
func connect2(left, right *Node) {
	left.Next = right
	if left.Right != nil {
		connect2(left.Left, left.Right)
		connect2(left.Right, right.Left)
		connect2(right.Left, right.Right)
	}
}

// @lc code=end

