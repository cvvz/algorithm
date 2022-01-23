/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	// 从i往上爬最少花多少钱
	// dp := make([]int, n, n)
	// dp[0] = cost[0]
	// dp[1] = cost[1]

	a := cost[0]
	b := cost[1]

	for i := 2; i < n; i++ {
		// dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
		a, b = b, min(a, b)+cost[i]
	}

	// return min(dp[n-1], dp[n-2])
	return min(a, b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

