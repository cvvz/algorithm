/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
type MedianFinder struct {
	maxHeap *MaxIntHeap
	minHeap *IntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: &MaxIntHeap{},
		minHeap: &IntHeap{},
	}
}

// 小顶堆
type IntHeap struct {
	sort.IntSlice
}

func (h *IntHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	v := h.IntSlice[h.IntSlice.Len()-1]
	h.IntSlice = h.IntSlice[:h.IntSlice.Len()-1]
	return v
}

func (h *IntHeap) Front() int {
	return h.IntSlice[0]
}

// 大顶堆, 官方答案用的是取负来实现大顶堆的效果
type MaxIntHeap struct {
	IntHeap
}

func (h *MaxIntHeap) Less(i, j int) bool {
	return h.IntHeap.IntSlice.Less(j, i)
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 || num < this.maxHeap.Front() {
		heap.Push(this.maxHeap, num)
		if this.maxHeap.Len() > this.minHeap.Len()+1 {
			heap.Push(this.minHeap, heap.Pop(this.maxHeap))
		}
	} else {
		heap.Push(this.minHeap, num)
		if this.minHeap.Len() > this.maxHeap.Len() {
			heap.Push(this.maxHeap, heap.Pop(this.minHeap))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(this.maxHeap.Front())
	} else {
		return float64(this.maxHeap.Front()+this.minHeap.Front()) / 2.0
	}
}

// 注意这里大顶堆的实现, 官方答案用的是取负来实现大顶堆的效果

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

