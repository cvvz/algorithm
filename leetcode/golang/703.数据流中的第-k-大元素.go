/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第 K 大元素
 */

// @lc code=start
type KthLargest struct {
	heap []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	minHeap := nums
	if len(nums) >= k {
		minHeap = nums[:k]
	}

	buildMinHeap(minHeap)

	for i := k; i < len(nums); i++ {
		if nums[i] > minHeap[0] {
			minHeap[0] = nums[i]
			minHeapify(minHeap, 0)
		}
	}

	return KthLargest{
		heap: minHeap,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	if len(this.heap) < this.k {
		this.heap = append(this.heap, val)
		minHeapify2(this.heap)
		return this.heap[0]
	}

	if val > this.heap[0] {
		this.heap[0] = val
		minHeapify(this.heap, 0)
	}

	return this.heap[0]
}

func buildMinHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		minHeapify(nums, i)
	}
}

// 从上往下堆化
func minHeapify(nums []int, i int) {
	min, l, r := i, i*2+1, i*2+2
	if l < len(nums) && nums[l] < nums[min] {
		min = l
	}
	if r < len(nums) && nums[r] < nums[min] {
		min = r
	}
	if min != i {
		nums[min], nums[i] = nums[i], nums[min]
		minHeapify(nums, min)
	}
}

// 从下往上堆化
func minHeapify2(nums []int) {
	min := len(nums) - 1
	for min > 0 {
		parent := min / 2
		if min%2 == 0 {
			parent = min/2 - 1
		}
		if nums[parent] <= nums[min] {
			break
		}

		nums[parent], nums[min] = nums[min], nums[parent]
		min = parent
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end

