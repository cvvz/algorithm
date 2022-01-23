/*
 * @lc app=leetcode.cn id=1822 lang=golang
 *
 * [1822] 数组元素积的符号
 */

// @lc code=start
func arraySign(nums []int) int {
	n := len(nums)
	result := 1

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			return 0
		} else if nums[i] < 0 {
			result = -result
		}
	}

	return result
}

// @lc code=end

