/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start

// 方法三：自建堆
// func findKthLargest(nums []int, k int) int {

// }

// 方法二： 快速选择算法
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

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

