/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 方法一：翻转链表，相加生成新链表，再次翻转链表
	// newL1 := reverseList(l1)
	// newL2 := reverseList(l2)
	// result := plusList(newL1, newL2)
	// return reverseList(result)

	// 方法二：存到栈中，然后依次相加
	stack1 := new(stack)
	stack2 := new(stack)

	for l1 != nil {
		stack1.Push(l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stack2.Push(l2.Val)
		l2 = l2.Next
	}

	dummyHead := &ListNode{}

	carry := 0

	for !stack1.empty() || !stack2.empty() {
		sum := carry
		if !stack1.empty() {
			sum += stack1.Pop()
		}
		if !stack2.empty() {
			sum += stack2.Pop()
		}

		carry = sum / 10
		sum = sum % 10
		node := &ListNode{
			Val:  sum,
			Next: dummyHead.Next,
		}
		dummyHead.Next = node
	}

	if carry != 0 {
		node := &ListNode{
			Val:  1,
			Next: dummyHead.Next,
		}
		dummyHead.Next = node
	}

	return dummyHead.Next
}

// 方法二
type stack []int

func (s *stack) Push(val int) {
	*s = append(*s, val)
}

func (s *stack) Pop() int {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

//////////////////////

// 方法一
func reverseList(head *ListNode) *ListNode {
	dummyHead := &ListNode{}

	for head != nil {
		next := head.Next
		head.Next = dummyHead.Next
		dummyHead.Next = head
		head = next
	}

	return dummyHead.Next
}

func plusList(l1, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tail := dummyHead

	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		node := &ListNode{
			Val: sum % 10,
		}
		tail.Next = node
		tail = tail.Next
	}

	if carry != 0 {
		node := &ListNode{
			Val: 1,
		}
		tail.Next = node
	}

	return dummyHead.Next
}

// @lc code=end

