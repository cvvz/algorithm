// https://leetcode-cn.com/problems/NYBBNL/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	inorderArr := inorder(root)
	rootAns := &TreeNode{
		Val: inorderArr[0],
	}

	p := rootAns
	for i := 1; i < len(inorderArr); i++ {
		p.Right = &TreeNode{
			Val: inorderArr[i],
		}
		p = p.Right
	}
	return rootAns
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	left := inorder(root.Left)
	ret = append(ret, left...)
	ret = append(ret, root.Val)
	return append(ret, inorder(root.Right)...)
}