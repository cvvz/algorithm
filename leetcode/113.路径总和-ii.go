/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	path := []int{}

	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}

		path = append(path, root.Val)
		defer func() {
			path = path[:len(path)-1]
		}()

		targetSum -= root.Val

		if targetSum == 0 && root.Left == nil && root.Right == nil {
			ans = append(ans, append([]int{}, path...))
			return
		}

		dfs(root.Left, targetSum)
		dfs(root.Right, targetSum)
	}

	dfs(root, targetSum)

	return ans
}

// @lc code=end

