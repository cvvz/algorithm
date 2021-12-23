// https://leetcode-cn.com/problems/UHnkqh/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	sentinelHead := &ListNode{}
	p := head

	for p != nil {
		q := p.Next
		p.Next = sentinelHead.Next
		sentinelHead.Next = p
		p = q
	}

	return sentinelHead.Next
}