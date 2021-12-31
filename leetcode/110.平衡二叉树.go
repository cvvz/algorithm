/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := getHeight(root.Left)
	r := getHeight(root.Right)
	if l-r < -1 || l-r > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)

}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := getHeight(root.Left)
	r := getHeight(root.Right)

	if l < r {
		return r + 1
	}
	return l + 1
}

// 时间复杂度：O(n2)
// 空间复杂度：O(n)

// @lc code=end

