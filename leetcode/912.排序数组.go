/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	// 快速排序
	// quickSort(nums, 0, len(nums)-1)

	// 🌟🌟🌟归并排序
	// mergeSort(nums, 0, len(nums)-1)

	return nums
}

func mergeSort(nums []int, l, r int) {
	if r <= l {
		return
	}

	mid := (l + r) / 2

	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)

	merge(nums, l, mid, r)
}

func merge(nums []int, l, mid, r int) {
	i, j := l, mid+1
	merged := []int{}

	for i <= mid && j <= r {
		if nums[i] < nums[j] {
			merged = append(merged, nums[i])
			i++
		} else {
			merged = append(merged, nums[j])
			j++
		}
	}

	for i <= mid {
		merged = append(merged, nums[i])
		i++
	}

	for j <= r {
		merged = append(merged, nums[j])
		j++
	}

	for i := 0; i < r-l+1; i++ {
		nums[l+i] = merged[i]
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
	nums[pivot], nums[end] = nums[end], nums[pivot]

	i := start

	for j := start; j < end; j++ {
		if nums[j] <= nums[end] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]

	return i
}

// @lc code=end

