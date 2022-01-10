/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return symmetric(root.Left, root.Right)
}

func symmetric(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 != nil && tree2 == nil || tree2 != nil && tree1 == nil {
		return false
	}

	if tree1.Val != tree2.Val {
		return false
	}

	return symmetric(tree1.Left, tree2.Right) && symmetric(tree2.Left, tree1.Right)
}

// 两颗树对称，意味着根节点相同，左子树和右子树是对称的
// @lc code=end

