/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parentMap := make(map[*TreeNode]*TreeNode)

	var preorderTraversal func(*TreeNode, *TreeNode)

	preorderTraversal = func(root, parent *TreeNode) {
		if root == nil {
			return
		}

		parentMap[root] = parent

		preorderTraversal(root.Left, root)
		preorderTraversal(root.Right, root)
	}

	preorderTraversal(root, nil)

	visited := make(map[*TreeNode]bool)

	for p != nil {
		visited[p] = true
		p = parentMap[p]
	}

	for q != nil {
		if visited[q] {
			return q
		}
		q = parentMap[q]
	}

	return nil
}

// 记录所有节点的父节点

// @lc code=end

