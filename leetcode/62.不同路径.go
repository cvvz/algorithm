/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([][]int, m, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n, n)
	}

	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j-1 >= 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]

}

// @lc code=end

