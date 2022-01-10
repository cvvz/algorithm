/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	currentQ := nodeQueue{root}

	for !currentQ.IsEmpty() {
		nextQ := nodeQueue{}
		currentResult := []int{}
		for !currentQ.IsEmpty() {
			top := currentQ.Pop()
			currentResult = append(currentResult, top.Val)
			if top.Left != nil {
				nextQ.Push(top.Left)
			}
			if top.Right != nil {
				nextQ.Push(top.Right)
			}
		}
		result = append(result, currentResult)
		currentQ = nextQ
	}

	return result
}

// BFS用队列实现
type nodeQueue []*TreeNode

func (q *nodeQueue) Push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *nodeQueue) Pop() *TreeNode {
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

func (q *nodeQueue) IsEmpty() bool {
	return len(*q) == 0
}

// 时间复杂度：O(n)，树的每个节点都遍历到了
// 空间复杂度：O(n)，队列的长度最长和二叉树的节点个树一样

// @lc code=end

