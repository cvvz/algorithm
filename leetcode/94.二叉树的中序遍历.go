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
	stack := new(nodeStack)
	ans := []int{}

	for root != nil || !stack.IsEmpty() {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}

		top := stack.Pop()
		ans = append(ans, top.Val)
		root = top.Right
	}

	return ans
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

// 沿着左子树一直遍历，如果当前节点不为空，入栈
// 出栈，右子树作为当前节点

// @lc code=end

