// https://leetcode-cn.com/problems/c32eOV/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// 1. 判断有没有环
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	fast, slow := head.Next.Next, head.Next
	for fast.Next != nil && fast.Next.Next != nil && fast != slow {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next == nil || fast.Next.Next == nil {
		return nil
	}

	// 2. 计算环的大小n，环的入口点就是整个链表的倒数第n个节点
	loopSize := 1
	for p := slow.Next; p != slow; p = p.Next {
		loopSize++
	}

	fast = head
	for i := 0; i < loopSize; i++ {
		fast = fast.Next
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}