/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start

// 构造栈的通用方法
type stack []int

func (this *stack) Push(val int) {
	*this = append(*this, val)
}

func (this *stack) Pop() {
	if len(*this) == 0 {
		return
	}
	*this = (*this)[:len(*this)-1]
}

func (this *stack) Top() int {
	return (*this)[len(*this)-1]
}

func (this *stack) Empty() bool {
	return len(*this) == 0
}

// 最小栈
type MinStack struct {
	valueStack stack
	minStack   stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		valueStack: stack{},
		minStack:   stack{},
	}
}

func (this *MinStack) Push(val int) {
	this.valueStack.Push(val)

	if this.minStack.Empty() || val < this.minStack.Top() {
		this.minStack.Push(val)
	} else {
		this.minStack.Push(this.minStack.Top())
	}
}

func (this *MinStack) Pop() {
	this.valueStack.Pop()
	this.minStack.Pop()
}

func (this *MinStack) Top() int {
	return this.valueStack.Top()
}

func (this *MinStack) GetMin() int {
	return this.minStack.Top()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

