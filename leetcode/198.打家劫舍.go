/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	// dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	// dp[0] = nums[0]
	// dp[1] = max(nums[0], nums[1])

	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n, n)

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

