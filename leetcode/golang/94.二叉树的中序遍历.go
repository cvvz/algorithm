/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	
	stack := []*TreeNode{root}
	ans := []int{}

	for len(stack) != 0 {
		top := stack[len(stack)-1]

		if top.Left != nil {
			stack = append(stack, top.Left)
			top.Left = nil
		} else {
			ans = append(ans, top.Val)
			stack = stack[:len(stack)-1]
			if top.Right != nil {
				stack = append(stack, top.Right)
				top.Right = nil
			}
		}
	}

	return ans
}

// 递归(递归怎么写代码量最少？)
// 时间复杂度：
// 空间复杂度：

// 迭代：
// 时间复杂度：
// 空间复杂度：
// @lc code=end

