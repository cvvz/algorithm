/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	// 快速排序
	// quickSort(nums, 0, len(nums)-1)
	// return nums

	// 归并排序
	ans := mergeSort(nums, 0, len(nums)-1)
	return ans
}

func mergeSort(nums []int, start, end int) []int {
	if start == end {
		return nums[start : end+1]
	}

	mid := (start + end) / 2

	sorted1 := mergeSort(nums, start, mid)
	sorted2 := mergeSort(nums, mid+1, end)
	sorted := merge(sorted1, sorted2)

	return sorted
}

func merge(sorted1, sorted2 []int) []int {
	var ret []int

	i, j := 0, 0
	for i < len(sorted1) && j < len(sorted2) {
		if sorted1[i] <= sorted2[j] {
			ret = append(ret, sorted1[i])
			i++
		} else {
			ret = append(ret, sorted2[j])
			j++
		}
	}

	ret = append(ret, sorted1[i:]...)
	ret = append(ret, sorted2[j:]...)

	return ret
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	i := partition(nums, start, end)

	quickSort(nums, start, i-1)
	quickSort(nums, i+1, end)
}

func partition(nums []int, start, end int) int {
	pivot := rand.Intn(end-start+1) + start
	nums[end], nums[pivot] = nums[pivot], nums[end]

	i := start
	for j := start; j < end; j++ {
		if nums[j] < nums[end] {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]

	return i
}

// @lc code=end

