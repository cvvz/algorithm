/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p!= nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left,q.Left) && isSameTree(p.Right, q.Right)
	} else if p == nil && q == nil {
		return true
	}

	return false
}

// 深度优先搜索
// @lc code=end

