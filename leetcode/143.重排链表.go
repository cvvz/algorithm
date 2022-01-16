/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	listArr := []*ListNode{}
	p := head

	for head != nil {
		listArr = append(listArr, head)
		head = head.Next
	}

	n := len(listArr)
	for i := n - 1; i >= 0; i-- {
		if listArr[i] == p || listArr[i] == p.Next {
			return
		}
		listArr[i].Next = p.Next
		p.Next = listArr[i]
		listArr[i-1].Next = nil

		p = p.Next.Next
	}
}

// @lc code=end

