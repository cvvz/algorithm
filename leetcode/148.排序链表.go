/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 找到链表的中点
	rListHead := getMidListNode(head)

	// 排序左右两个链表
	llist := sortList(head)
	rlist := sortList(rListHead)

	// 合并左右两个链表
	sortedList := mergeList(llist, rlist)

	return sortedList
}

func getMidListNode(head *ListNode) *ListNode {
	slow, quick := head, head

	for quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	rListHead := slow.Next
	slow.Next = nil
	return rListHead
}

func mergeList(lhead *ListNode, rhead *ListNode) *ListNode {
	var head, tail *ListNode

	for lhead != nil && rhead != nil {
		if lhead.Val <= rhead.Val {
			if head == nil {
				head = lhead
				tail = lhead
			} else {
				tail.Next = lhead
				tail = lhead
			}
			lhead = lhead.Next
		} else {
			if head == nil {
				head = rhead
				tail = rhead
			} else {
				tail.Next = rhead
				tail = rhead
			}
			rhead = rhead.Next
		}
	}

	if lhead != nil {
		tail.Next = lhead
	}
	if rhead != nil {
		tail.Next = rhead
	}

	return head
}

// @lc code=end

