/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 */

// @lc code=start
type MyLinkedList struct {
	Head *MyListNode
	Tail *MyListNode
	Len  int
}

type MyListNode struct {
	Val  int
	Next *MyListNode
	Prev *MyListNode
}

func Constructor() MyLinkedList {
	dummyHead, dummyTail := &MyListNode{}, &MyListNode{}
	dummyHead.Next = dummyTail
	dummyTail.Prev = dummyHead

	return MyLinkedList{
		Head: dummyHead,
		Tail: dummyTail,
		Len:  0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Len {
		return -1
	}

	head := this.Head.Next

	for i := 0; i < index; i++ {
		head = head.Next
	}

	return head.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &MyListNode{
		Val:  val,
		Next: this.Head.Next,
		Prev: this.Head,
	}
	this.Head.Next.Prev = node
	this.Head.Next = node

	this.Len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &MyListNode{
		Val:  val,
		Next: this.Tail,
		Prev: this.Tail.Prev,
	}
	this.Tail.Prev.Next = node
	this.Tail.Prev = node

	this.Len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Len {
		return
	} else if index < 0 {
		this.AddAtHead(val)
	} else if index == this.Len {
		this.AddAtTail(val)
	} else {
		head := this.Head.Next
		for i := 0; i < index; i++ {
			head = head.Next
		}

		node := &MyListNode{
			Val: val,
		}
		prev := head.Prev

		node.Next = head
		head.Prev = node
		prev.Next = node
		node.Prev = prev

		this.Len++
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Len {
		return
	}

	head := this.Head.Next

	for i := 0; i < index; i++ {
		head = head.Next
	}

	prev := head.Prev
	next := head.Next

	prev.Next = next
	next.Prev = prev

	this.Len--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end

