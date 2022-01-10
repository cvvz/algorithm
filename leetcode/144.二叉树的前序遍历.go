/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}

	stack := new(nodeStack)

	for root != nil || !stack.IsEmpty() {
		for root != nil {
			ans = append(ans, root.Val)
			stack.Push(root)
			root = root.Left
		}

		top := stack.Pop()
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

// 若根节点非空，从根节点出发，沿着左子树，将当前节点入栈
// 出栈，根节点变为右节点
// @lc code=end

