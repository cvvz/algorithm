/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := len(nums)

	for left <= right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

// 二分法

// 时间复杂度：O(logn)
// 空间复杂度：O(1)
// @lc code=end

