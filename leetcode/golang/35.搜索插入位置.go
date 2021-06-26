/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	first, last := 0, len(nums)-1

	for first < last {
		mid := (first + last) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}

	if target > nums[first] {
		return first + 1
	} else {
		return first
	}
}

// 时间复杂度：O(logn)
// 空间复杂度：O(1)
// @lc code=end

