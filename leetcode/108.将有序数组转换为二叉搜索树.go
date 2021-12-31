/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	mid := nums[length/2]
	leftNums := nums[:length/2]
	rightNums := nums[length/2+1:]

	root := &TreeNode{
		Val:   mid,
		Left:  sortedArrayToBST(leftNums),
		Right: sortedArrayToBST(rightNums),
	}

	return root
}

// 二叉搜索树：左<根<右
// 中序遍历（左根右）结果转成二叉树
// 为了高度平衡：没有左子树或者没有右子树

// 时间复杂度：O(n)
// 空间复杂度：O(logn)
// @lc code=end

