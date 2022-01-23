/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := make([][]int, m, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n, n)
	}

	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 && obstacleGrid[i-1][j] != 1 {
				dp[i][j] += dp[i-1][j]
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] != 1 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

// @lc code=end

