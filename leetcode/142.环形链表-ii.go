/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	// 快慢指针判断环
	quick, slow := head, head

	for quick != nil && quick.Next != nil && quick.Next.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
		if quick == slow {
			break
		}
	}

	if quick != slow {
		return nil
	}

	// 绕环一周获得环的大小n
	circleSize := 1
	p := slow.Next

	for p != slow {
		circleSize++
		p = p.Next
	}

	// 环的入口即倒数第n个节点
	p, q := head, head
	for i := 0; i < circleSize; i++ {
		p = p.Next
	}

	for p != q {
		p = p.Next
		q = q.Next
	}

	return q
}

// @lc code=end

