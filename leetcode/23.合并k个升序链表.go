/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummyHead := &ListNode{}
	tail := dummyHead

	for {
		headList := []*ListNode{}
		for _, head := range lists {
			headList = append(headList, head)
		}

		i := func() int {
			index := 0
			min := 100000

			for i, head := range headList {
				if head != nil && head.Val < min {
					index = i
					min = head.Val
				}
			}

			return index
		}()

		if lists[i] == nil {
			break
		}

		tail.Next = lists[i]
		tail = tail.Next

		lists[i] = lists[i].Next
	}

	return dummyHead.Next
}

// @lc code=end

