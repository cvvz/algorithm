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
	var ans []int
	var s stack

	if root == nil {
		return ans
	}

	s.push(root)

	for !s.empty() {
		node := s.top()
		if node.Left != nil {
			s.push(node.Left)
			node.Left = nil
		} else {
			_ = s.pop()
			ans = append(ans, node.Val)
			if node.Right != nil {
				s.push(node.Right)
				node.Right = nil
			}
		}
	}

	return ans
}

type stack []*TreeNode

func (s *stack) push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *stack) pop() *TreeNode {
	top := s.top()
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *stack) top() *TreeNode {
	return (*s)[len(*s)-1]
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

//🌟 二叉查找树用中序遍历能得到顺序数组
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// @lc code=end

