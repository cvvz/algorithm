/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
func rob(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 偷root
	total := root.Val
	if root.Left != nil {
		total = total + dfs(root.Left.Left) + dfs(root.Left.Right)
	}
	if root.Right != nil {
		total = total + dfs(root.Right.Left) + dfs(root.Right.Right)
	}

	max := total

	// 不偷root
	total = dfs(root.Left) + dfs(root.Right)

	if max > total {
		return max
	}

	return total
}

// @lc code=end

