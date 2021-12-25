// https://leetcode-cn.com/problems/lMSNwu/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)

	sentinelHead := &ListNode{}
	p, q := l1, l2
	flag := 0
	var sum int

	for p != nil || q != nil {
		if p == nil {
			sum = flag + q.Val
			q = q.Next
		} else if q == nil {
			sum = flag + p.Val
			p = p.Next
		} else {
			sum = flag + p.Val + q.Val
			q = q.Next
			p = p.Next
		}

		node := &ListNode{
			Val: sum % 10,
		}
		flag = sum / 10

		insertHead(sentinelHead, node)
	}

	if flag != 0 {
		node := &ListNode{
			Val: 1,
		}

		insertHead(sentinelHead, node)
	}

	return sentinelHead.Next
}

func insertHead(head *ListNode, node *ListNode) {
	node.Next = head.Next
	head.Next = node
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