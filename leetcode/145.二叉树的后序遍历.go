/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {

}

type nodeStack []*TreeNode

func (s *nodeStack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *nodeStack) Pop() *TreeNode {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *nodeStack) IsEmpty() bool {
	return len(*s) == 0
}

// 如果当前节点不为空，从当前节点出发沿着左子树遍历
// 右子树变为当前节点

// @lc code=end

