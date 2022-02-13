/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDiameter := 0

	// 求最大深度的过程中，顺便求最大直径
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		diameter := 0
		var lDepth, rDepth int

		if root.Left != nil {
			lDepth = dfs(root.Left) + 1
			diameter += lDepth
		}

		if root.Right != nil {
			rDepth = dfs(root.Right) + 1
			diameter += rDepth
		}

		if diameter > maxDiameter {
			maxDiameter = diameter
		}

		return max(lDepth, rDepth)
	}

	dfs(root)

	return maxDiameter
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

