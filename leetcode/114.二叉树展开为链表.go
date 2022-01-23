/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	preOrderTraversal(root)
}

func preOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	right := root.Right

	root.Right = root.Left

	preOrderTraversal(root.Left)

	root.Left = nil

	for root.Right != nil {
		root = root.Right
	}

	root.Right = right

	preOrderTraversal(root.Right)
}

// @lc code=end

