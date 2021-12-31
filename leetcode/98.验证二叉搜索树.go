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
	bstArr := inorderTraversal(root)

	for i := 1; i < len(bstArr); i++ {
		if bstArr[i] <= bstArr[i-1] {
			return false
		}
	}
	return true
}

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var foo func(*TreeNode)

	foo = func(root *TreeNode) {
		if root == nil {
			return
		}
		foo(root.Left)
		ans = append(ans, root.Val)
		foo(root.Right)
	}

	foo(root)
	return ans
}

// 中序遍历一个二叉搜索树能得到一个从小到大排列的数组

// 时间复杂度：O(n)
// 空间复杂度：O(n)
// @lc code=end

