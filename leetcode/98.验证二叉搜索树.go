/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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

func isValidBST(root *TreeNode) bool {
	result := inorderTraversal(root)
	n := len(result)
	for i := 1; i < n; i++ {
		if result[i] <= result[i-1] {
			return false
		}
	}

	return true
}

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}

	stack := new(nodeStack)

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

// 中序遍历一个二叉搜索树能得到一个从小到大排列的数组

// 时间复杂度：O(n)
// 空间复杂度：O(n)
// @lc code=end

