/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	old2new := make(map[*Node]*Node)
	sentinel := &Node{}
	tail := sentinel

	for head != nil {
		newNode := &Node{
			Val: head.Val,
		}
		tail.Next = newNode
		tail = tail.Next
		old2new[head] = newNode
		head = head.Next
	}

	for oldNode, newNode := range old2new {
		newNode.Random = old2new[oldNode.Random]
	}

	return sentinel.Next
}

// 使用map[node]newnode存储新老节点的映射

// @lc code=end

