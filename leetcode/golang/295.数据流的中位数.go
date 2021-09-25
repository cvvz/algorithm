/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start

/****** 不使用库函数，底层实现细节和库函数是一样的 *******/

type MedianFinder struct {
	minHeap MinHeap
	maxHeap MaxHeap
}

type Interface interface {
	Less(int, int) bool
	Swap(int, int)
	Len() int
}

type MinHeap []int
type MaxHeap []int

func (mh MinHeap) Less(i, j int) bool {
	return mh[i] < mh[j]
}

func (mh MaxHeap) Less(i, j int) bool {
	return mh[i] > mh[j]
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MaxHeap) Len() int {
	return len(mh)
}

func HeapifyUp(h Interface, i int) {
	for {
		parent := (i - 1) / 2
		if i == parent || !h.Less(i, parent) {
			break
		}
		h.Swap(i, parent)
		i = parent
	}
}

func HeapifyDown(h Interface, i int) {
	parent := i
	l := i*2 + 1
	r := i*2 + 2

	if l < h.Len() && h.Less(l, parent) {
		parent = l
	}
	if r < h.Len() && h.Less(r, parent) {
		parent = r
	}
	if parent != i {
		h.Swap(i, parent)
		HeapifyDown(h, parent)
	}
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: MinHeap{},
		maxHeap: make(MaxHeap, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.maxHeap) == 0 {
		this.maxHeap = append(this.maxHeap, num)
		return
	}
	if num <= this.maxHeap[0] {
		if len(this.maxHeap) > len(this.minHeap) {
			this.minHeap = append(this.minHeap, this.maxHeap[0])
			HeapifyUp(this.minHeap, len(this.minHeap)-1)
			this.maxHeap[0] = num
			HeapifyDown(this.maxHeap, 0)
		} else {
			this.maxHeap = append(this.maxHeap, num)
			HeapifyUp(this.maxHeap, len(this.maxHeap)-1)
		}
	} else {
		if len(this.maxHeap) > len(this.minHeap) {
			this.minHeap = append(this.minHeap, num)
			HeapifyUp(this.minHeap, len(this.minHeap)-1)
		} else {
			this.minHeap = append(this.minHeap, num)
			HeapifyUp(this.minHeap, len(this.minHeap)-1)

			this.maxHeap = append(this.maxHeap, this.minHeap[0])
			HeapifyUp(this.maxHeap, len(this.maxHeap)-1)

			this.minHeap[0] = this.minHeap[len(this.minHeap)-1]
			this.minHeap = this.minHeap[:len(this.minHeap)-1]
			HeapifyDown(this.minHeap, 0)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxHeap) > len(this.minHeap) {
		return float64(this.maxHeap[0])
	}
	return float64((this.minHeap[0] + this.maxHeap[0])) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

