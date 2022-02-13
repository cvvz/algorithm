// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return isContain(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isContain(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil || A.Val != B.Val {
		return false
	}

	return isContain(A.Left, B.Left) && isContain(A.Right, B.Right)
}