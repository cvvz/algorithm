/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			p.Next = l2
			return head.Next
		}
		if l2 == nil {
			p.Next = l1
			return head.Next
		}
		l := l2
		if l1.Val < l2.Val {
			l = l1
			l1 = l1.Next
		} else {
			l2 = l2.Next
		}
		p.Next, p, l = l, l, l.Next
	}
	return nil
}

// 时间复杂度：O(m+n)
// 空间复杂度：O(1)

// @lc code=end

