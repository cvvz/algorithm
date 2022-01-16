/*
 * @lc app=leetcode.cn id=1448 lang=golang
 *
 * [1448] 统计二叉树中好节点的数目
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
func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var count int

	var preorderTraversal func(*TreeNode, int)

	preorderTraversal = func(root *TreeNode, max int) {
		if root == nil {
			return
		}

		if root.Val >= max {
			count++
			max = root.Val
		}

		preorderTraversal(root.Left, max)
		preorderTraversal(root.Right, max)
	}

	preorderTraversal(root, root.Val)

	return count

}

// @lc code=end

