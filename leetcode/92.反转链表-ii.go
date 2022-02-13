/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}

	prev, cur := dummyHead, head

	for i := 1; i < left; i++ {
		prev = cur
		cur = cur.Next
	}

	dummyHead2 := &ListNode{}
	for i := left; i <= right; i++ {
		next := cur.Next
		cur.Next = dummyHead2.Next
		dummyHead2.Next = cur
		cur = next
	}

	prev.Next = dummyHead2.Next

	for prev.Next != nil {
		prev = prev.Next
	}
	prev.Next = cur

	return dummyHead.Next
}

// @lc code=end

