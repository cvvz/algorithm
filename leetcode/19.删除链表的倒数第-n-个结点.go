/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{0, head}
	prev, p := dummyHead, dummyHead

	for i := 0; i < n; i++ {
		p = p.Next
	}

	for p.Next != nil {
		p = p.Next
		prev = prev.Next
	}

	prev.Next = prev.Next.Next

	return dummyHead.Next
}

// @lc code=end

