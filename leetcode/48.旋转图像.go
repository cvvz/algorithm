/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start

// 借助额外的数组
func rotate(matrix [][]int) {
	n := len(matrix)

	temp := make([][]int, n)
	for i := 0; i < n; i++ {
		temp[i] = make([]int, n)
	}

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			tr := c
			tc := n - 1 - r
			temp[tr][tc] = matrix[r][c]
		}
	}

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			matrix[r][c] = temp[r][c]
		}
	}
}

// @lc code=end

