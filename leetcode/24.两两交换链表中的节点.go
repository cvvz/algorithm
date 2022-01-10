/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	sentinel := &ListNode{
		Next: head,
	}
	prev := sentinel
	p := head

	for p != nil && p.Next != nil {
		q := p.Next

		prev.Next = q
		p.Next = q.Next
		q.Next = p

		prev = p
		p = p.Next
	}

	return sentinel.Next
}

// @lc code=end

