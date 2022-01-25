/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p, q := l1, l2
	carry := 0
	dummyHead := &ListNode{}
	tail := dummyHead

	for p != nil || q != nil {
		sum := 0
		if p != nil {
			sum += p.Val
			p = p.Next
		}
		if q != nil {
			sum += q.Val
			q = q.Next
		}
		sum += carry
		carry = sum / 10
		tail.Next = &ListNode{
			Val: sum % 10,
		}
		tail = tail.Next
	}
	if carry == 1 {
		tail.Next = &ListNode{
			Val: 1,
		}
	}
	return dummyHead.Next
}

// @lc code=end

