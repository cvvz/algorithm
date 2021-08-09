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
	ans := &ListNode{}
	k := ans

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
		k.Next = &ListNode{
			Val: sum % 10,
		}
		k = k.Next
	}
	if carry == 1 {
		k.Next = &ListNode{
			Val: 1,
		}
	}
	return ans.Next
}

// @lc code=end

