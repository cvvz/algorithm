// https://leetcode-cn.com/problems/aMhZSa/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	p := head
	count := 0

	for p != nil {
		count++
		p = p.Next
	}

	p = head
	// 翻转后半部分
	for i := 0; i < count/2; i++ {
		p = p.Next
	}

	p = reverse(p)
	// 比较
	q := head
	for p != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
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