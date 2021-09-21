/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	// 堆排序
	// 先利用前k个元素建小顶堆
	minHeap := nums[:k]
	buildMinHeap(minHeap)
	// 如果比堆顶大，那么就替换堆顶元素，然后进行一次堆化
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			minHeapify(minHeap, 0)
		}
	}
	// 返回堆顶元素
	return nums[0]
}

// 建堆
func buildMinHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		minHeapify(nums, i)
	}
}

// 🌟🌟堆化
func minHeapify(nums []int, i int) {
	l := i*2 + 1
	r := i*2 + 2
	min := i

	if l < len(nums) && nums[l] < nums[min] {
		min = l
	}
	if r < len(nums) && nums[r] < nums[min] {
		min = r
	}
	if min != i {
		nums[i], nums[min] = nums[min], nums[i]
		minHeapify(nums, min)
	}
}

// @lc code=end

