/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
type stack []*TreeNode

func (s *stack) push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *stack) pop() *TreeNode {
	node := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return node
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var s1, s2 stack
	var ans [][]int
	odd := true

	if root == nil {
		return ans
	}

	s1.push(root)

	for !s1.empty() {
		levelArr := []int{}

		for !s1.empty() {
			node := s1.pop()
			levelArr = append(levelArr, node.Val)
			if odd {
				if node.Left != nil {
					s2.push(node.Left)
				}
				if node.Right != nil {
					s2.push(node.Right)
				}
			} else {
				if node.Right != nil {
					s2.push(node.Right)
				}
				if node.Left != nil {
					s2.push(node.Left)
				}
			}
		}

		odd = !odd
		ans = append(ans, levelArr)
		s1 = s2
		s2 = stack{}
	}

	return ans
}

// 🌟这里层序遍历用到的是栈，而不是队列
// 奇数层从左往右入栈
// 偶数层从右往左入栈

// @lc code=end

