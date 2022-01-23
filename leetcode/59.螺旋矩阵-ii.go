/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	ans := make([][]int, n, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n, n)
	}

	total := n * n

	// 右下左上
	directionSet := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	i, j := 0, 0
	directionIndex := 0

	for count := 1; count <= total; count++ {
		ans[i][j] = count

		direction := directionSet[directionIndex]
		x := direction[0]
		y := direction[1]

		if i+x < 0 || i+x >= n || j+y < 0 || j+y >= n || ans[i+x][j+y] != 0 {
			// 转向
			directionIndex = (directionIndex + 1) % 4
			direction = directionSet[directionIndex]
			x = direction[0]
			y = direction[1]
		}

		i, j = i+x, j+y
	}

	return ans
}

// @lc code=end

