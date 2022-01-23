/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	// 归并排序
	// mergeSort(nums, 0, len(nums)-1)

	// 快速排序
	// quickSort(nums, 0, len(nums)-1)

	// 堆排序
	heapSort(nums)

	return nums
}

// 2022.01.16

func heapSort(nums []int) {
	heap := new(MyHeap)
	n := len(nums)

	// 建堆
	for i := 0; i < n; i++ {
		heap.push(nums[i])
	}

	// 不断的pop
	for i := 0; i < n; i++ {
		nums[i] = heap.pop()
	}
}

type MyHeap []int

func (h *MyHeap) push(val int) {
	*h = append(*h, val)
	up(*h, len(*h)-1)
}

func (h *MyHeap) pop() int {
	top := (*h)[0]
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]
	*h = (*h)[:len(*h)-1]
	down(*h, 0)
	return top
}

func up(h MyHeap, i int) {
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

func down(h MyHeap, i int) {
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

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	pivot := partition(nums, start, end)

	quickSort(nums, start, pivot-1)
	quickSort(nums, pivot+1, end)
}

func partition(nums []int, start, end int) int {
	pivot := rand.Intn(end-start) + start

	nums[end], nums[pivot] = nums[pivot], nums[end]

	i := start

	for j := start; j < end; j++ {
		if nums[j] < nums[end] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[end] = nums[end], nums[i]

	return i
}

func mergeSort(nums []int, start, end int) {
	if start == end {
		return
	}
	mid := (start + end) / 2
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)
	merge(nums, start, end)
}

func merge(nums []int, start, end int) {
	mid := (start + end) / 2

	i, j := start, mid+1
	temp := []int{}
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}
	if i > mid {
		temp = append(temp, nums[j:end+1]...)
	}
	if j > end {
		temp = append(temp, nums[i:mid+1]...)
	}

	for i := 0; i < end-start+1; i++ {
		nums[start+i] = temp[i]
	}
}

// @lc code=end

