/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	ans := [][]int{}

	for i := 0; len(queue) != 0; i++ {
		// 浅拷贝，共用底层的数组。不写底层的数组就没问题
		tempQ := queue
		queue = []*TreeNode{}
		// 🌟这一行不能少！！！初始化二维数组为空（与是不是nil没有关系）时，访问下标0会越界🌟
		ans = append(ans, []int{})
		for len(tempQ) != 0 {
			ans[i] = append(ans[i], tempQ[0].Val)
			if tempQ[0].Left != nil {
				queue = append(queue, tempQ[0].Left)
			}
			if tempQ[0].Right != nil {
				queue = append(queue, tempQ[0].Right)
			}
			tempQ = tempQ[1:]
		}
	}
	return ans
}

// BFS用队列实现
// ❕❕slice的使用——如果没有规定长度或者初始化，访问元素会range out：
// var a []int // or a := []int{}
// fmt.Println(a) // []
// a[0] = 1 // 报错

// ❕❕注意，vscode中compile error报错信息中，
// 起始 Line 是@lc code=start这一行的下一行

// 时间复杂度：O(n)，树的每个节点都遍历到了
// 空间复杂度：O(n)，队列的长度最长和二叉树的节点个树一样

// @lc code=end

