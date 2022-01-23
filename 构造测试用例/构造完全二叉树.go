// 1.给定一组整数，请按照从上到下，从左到右的顺序构造一棵二叉树(其实就是完全二叉树)
// 2.对这个完全二叉树进行层次遍历

package main

import "fmt"

func main() {
	root := construct([]int{-1, 3, 0, 1, 2, 100, 4, 5, 6, 7, 8, 9, 10})
	levelOrderTraversal(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func construct(nums []int) *TreeNode {
	n := len(nums)

	if n == 0 {
		return nil
	}

	queue := new(nodeQueue)
	nextQueue := new(nodeQueue)

	root := &TreeNode{
		Val: nums[0],
	}
	queue.Push(root)

	i := 1
	for !queue.IsEmpty() {
		for !queue.IsEmpty() {
			node := queue.Pop()
			if i >= n {
				return root
			}
			node.Left = &TreeNode{
				Val: nums[i],
			}
			nextQueue.Push(node.Left)
			i++
			if i >= n {
				return root
			}
			node.Right = &TreeNode{
				Val: nums[i],
			}
			nextQueue.Push(node.Right)
			i++
		}
		queue = nextQueue
		nextQueue = new(nodeQueue)
	}

	return root
}

func levelOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	currentQ, nextQ := new(nodeQueue), new(nodeQueue)

	currentQ.Push(root)

	for !currentQ.IsEmpty() {
		for !currentQ.IsEmpty() {
			node := currentQ.Pop()
			fmt.Printf("%d ", node.Val)
			if node.Left != nil {
				nextQ.Push(node.Left)
			}
			if node.Right != nil {
				nextQ.Push(node.Right)
			}
		}
		currentQ = nextQ
		nextQ = new(nodeQueue)
	}
}

type nodeQueue []*TreeNode

func (q *nodeQueue) Push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *nodeQueue) Pop() *TreeNode {
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

func (q *nodeQueue) IsEmpty() bool {
	return len(*q) == 0
}
