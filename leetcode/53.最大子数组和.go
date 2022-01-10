/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	sumSubArray := nums[0]
	ans := sumSubArray

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for i := 1; i < len(nums); i++ {
		sumSubArray = max(sumSubArray+nums[i], nums[i])
		if sumSubArray > ans {
			ans = sumSubArray
		}
	}

	return ans
}

// DP，最重要是要找到推导公式
// @lc code=end

