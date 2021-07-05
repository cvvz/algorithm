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
	
	var twoTreeSymmetric func(left *TreeNode, right *TreeNode) bool

	twoTreeSymmetric = func(t1 *TreeNode, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		} 

		if t1 != nil && t2 != nil && t1.Val == t2.Val && twoTreeSymmetric(t1.Left, t2.Right) && twoTreeSymmetric(t1.Right, t2.Left) {
			return true
		}

		return false
	}

	return twoTreeSymmetric(root.Left, root.Right)
}


// 两棵数怎样才算是镜像对称：
// 根节点相同
// 左子树与右子树镜像对称
// @lc code=end

