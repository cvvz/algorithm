// https://leetcode-cn.com/problems/NaqhDT/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	root *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	return CBTInserter{
		root: root,
	}
}

func (this *CBTInserter) Insert(v int) int {
	// 层序遍历，找到第一个child为nil的节点
	var q1 queue
	q1.push(this.Get_root())

	for !q1.empty() {
		node := q1.pop()
		if node.Left == nil {
			node.Left = &TreeNode{
				Val: v,
			}
			return node.Val
		}
		if node.Right == nil {
			node.Right = &TreeNode{
				Val: v,
			}
			return node.Val
		}
		q1.push(node.Left)
		q1.push(node.Right)
	}

	return 0
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

type queue []*TreeNode

func (q *queue) push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *queue) pop() *TreeNode {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */