/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start

// 方法三：自建小顶堆
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	h := &myheap{}

	i := 0

	for i < k && i < n {
		h.push(nums[i])
		i++
	}

	for i < n {
		if nums[i] > (*h)[0] {
			h.push(nums[i])
			h.pop()
		}
		i++
	}

	return (*h)[0]
}

type myheap []int

func (h *myheap) push(val int) {
	*h = append(*h, val)
	up(*h, len(*h)-1)
}

func (h *myheap) pop() {
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	*h = (*h)[:len(*h)-1]
	down(*h, 0)
}

func up(h myheap, i int) {
	for i > 0 {
		//🌟 这里不需要判断是否比兄弟节点小
		parent := (i - 1) >> 1

		if i == parent || h[i] >= h[parent] {
			break
		}

		h[i], h[parent] = h[parent], h[i]
		i = parent
	}
}

func down(h myheap, i int) {
	for i < len(h) {
		left, right := i<<1+1, i<<1+2
		if left > len(h)-1 {
			return
		}

		toSwap := left
		if right <= len(h)-1 && h[left] > h[right] {
			toSwap = right
		}

		if h[toSwap] < h[i] {
			h[toSwap], h[i] = h[i], h[toSwap]
		}

		i = toSwap
	}

}

// 方法二： 快速选择算法
// func findKthLargest(nums []int, k int) int {
// 	return quickSelect(nums, 0, len(nums)-1, k)
// }

func quickSelect(nums []int, start, end, k int) int {
	pivot := partition(nums, start, end)
	if pivot == k-1 {
		return nums[pivot]
	}

	if pivot < k-1 {
		return quickSelect(nums, pivot+1, end, k)
	}

	return quickSelect(nums, start, pivot-1, k)
}

func partition(nums []int, start, end int) int {
	if start == end {
		return start
	}

	pivot := rand.Intn(end-start) + start

	nums[end], nums[pivot] = nums[pivot], nums[end]

	i := start

	for j := start; j < end; j++ {
		if nums[j] > nums[end] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[end] = nums[end], nums[i]

	return i
}

// 方法一： 使用heap官方库
// func findKthLargest(nums []int, k int) int {
// 	ih := new(IntHeap)

// 	// 构造大小为K的小顶堆
// 	i := 0
// 	for i < k && i < len(nums) {
// 		heap.Push(ih, nums[i])
// 		i++
// 	}

// 	// 依次和堆顶元素进行比较，比它大则Push，然后Pop
// 	for i < len(nums) {
// 		if nums[i] > ih.Top().(int) {
// 			heap.Push(ih, nums[i])
// 			_ = heap.Pop(ih)
// 		}
// 		i++
// 	}

// 	return ih.Top().(int)
// }

type IntHeap struct {
	sort.IntSlice
}

// add x as element Len()
func (h *IntHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

// remove and return element Len() - 1.
func (h *IntHeap) Pop() interface{} {
	top := h.IntSlice[h.IntSlice.Len()-1]
	h.IntSlice = h.IntSlice[:h.IntSlice.Len()-1]
	return top
}

// top of heap
func (h *IntHeap) Top() interface{} {
	return h.IntSlice[0]
}

// @lc code=end

