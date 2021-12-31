// https://leetcode-cn.com/problems/SLwz0R/submissions/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentinelHead := &ListNode{
		Next: head,
	}

	first, last, prev := head, head, sentinelHead

	for i := 0; i < n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		prev = last
		last = last.Next
	}

	prev.Next = last.Next

	return sentinelHead.Next
}