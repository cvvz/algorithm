/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 获得链表长度
	lengthA, lengthB := 0, 0
	ha, hb := headA, headB
	for ha != nil {
		ha = ha.Next
		lengthA++
	}

	for hb != nil {
		hb = hb.Next
		lengthB++
	}
	// 移动长链表的头和短链表对齐
	if lengthA > lengthB {
		for i := 0; lengthB+i < lengthA; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; lengthA+i < lengthB; i++ {
			headB = headB.Next
		}
	}
	// 同时移动直到链表尾，判断是否相同
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

// @lc code=end

