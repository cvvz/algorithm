/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	
	pre := head
	p := pre.Next

	for p != nil {
		if p.Val == pre.Val {
			pre.Next = p.Next
			p = p.Next
		} else {
			pre = pre.Next
			p = p.Next
		}
	}

	return head
}
// @lc code=end

