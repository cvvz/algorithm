/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
type queue []*TreeNode

func (q *queue) push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *queue) pop() *TreeNode {
	node := (*q)[0]
	*q = (*q)[1:]
	return node
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func levelOrderBottom(root *TreeNode) [][]int {
	var q1, q2 queue
	ans := [][]int{}

	if root == nil {
		return ans
	}

	q1.push(root)

	for !q1.isEmpty() {
		subAns := []int{}
		for !q1.isEmpty() {
			node := q1.pop()
			subAns = append(subAns, node.Val)
			if node.Left != nil {
				q2.push(node.Left)
			}
			if node.Right != nil {
				q2.push(node.Right)
			}
		}
		ans = append(ans, subAns)
		q1 = q2
		q2 = queue{}
	}

	reverse := [][]int{}
	for i := len(ans) - 1; i >= 0; i-- {
		reverse = append(reverse, ans[i])
	}

	return reverse
}

// @lc code=end

