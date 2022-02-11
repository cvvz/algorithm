/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, n, n)

	// 第一个房间跳过的情况得到的dp结果
	dp[1] = nums[1]
	dp[2] = max(nums[1], nums[2])

	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	// 最后一个房间跳过的情况得到的dp结果
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n-1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return max(dp[n-1], dp[n-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

