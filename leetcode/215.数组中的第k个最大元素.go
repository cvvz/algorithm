/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, begin int, end int, k int) int {
	i := partition(nums, begin, end)
	if i == k-1 {
		return nums[i]
	} else if i < k-1 {
		return quickSelect(nums, i+1, end, k)
	} else {
		return quickSelect(nums, begin, i-1, k)
	}
}

func partition(nums []int, begin int, end int) int {
	i := begin
	for j := begin; j < end; j++ {
		if nums[j] > nums[end] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[end] = nums[end], nums[i]
	return i
}

// 🌟🌟🌟快速排序的关键是partition函数

// @lc code=end

