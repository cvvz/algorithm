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
	set := make(map[*ListNode]struct{})

	for p := headA; p != nil; p = p.Next {
		set[p] = struct{}{}
	}

	for p := headB; p != nil; p = p.Next {
		if _, exist := set[p]; exist {
			return p
		}
	}

	return nil

}

// @lc code=end

