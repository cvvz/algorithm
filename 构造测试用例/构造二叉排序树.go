// 1. 给定一组整数，请构造一棵二叉排序树
// 2. 对二叉排序树进行前序遍历
package main

import "fmt"

func main() {
	root := construct([]int{2, 3, 1, 8, 10, 5, 6, 9, 7, 0, 4, -1, -4})
	preorderTraversal(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一组整数，请构造一棵二叉排序树
func construct(arr []int) *TreeNode {
	n := len(arr)
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val: arr[0],
	}
	for i := 1; i < n; i++ {
		insert(root, arr[i])
	}

	return root
}

func insert(root *TreeNode, val int) {
	var parent *TreeNode
	var isLeft bool

	for root != nil {
		parent = root
		if root.Val > val {
			isLeft = true
			root = root.Left
		} else {
			isLeft = false
			root = root.Right
		}
	}

	node := &TreeNode{
		Val: val,
	}

	if isLeft {
		parent.Left = node
	} else {
		parent.Right = node
	}

}

// 对二叉排序树进行前序遍历
func preorderTraversal(root *TreeNode) {
	stack := new(nodeStack)

	for !stack.IsEmpty() || root != nil {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}

		top := stack.Pop()
		fmt.Println(top.Val)
		root = top.Right
	}
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
