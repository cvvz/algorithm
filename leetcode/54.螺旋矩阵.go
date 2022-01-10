/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	//-100 <= matrix[i][j] <= 100
	visited := -101

	row := len(matrix)
	column := len(matrix[0])

	// 🌟定义边界
	top, bottom, left, right := 0, row-1, 0, column-1

	for len(ans) < row*column {
		for ci := left; ci <= right && matrix[top][ci] != visited; ci++ {
			ans = append(ans, matrix[top][ci])
			matrix[top][ci] = visited
		}
		for ri := top + 1; ri <= bottom && matrix[ri][right] != visited; ri++ {
			ans = append(ans, matrix[ri][right])
			matrix[ri][right] = visited
		}
		for ci := right - 1; ci >= left && matrix[bottom][ci] != visited; ci-- {
			ans = append(ans, matrix[bottom][ci])
			matrix[bottom][ci] = visited
		}
		for ri := bottom - 1; ri >= top+1 && matrix[ri][left] != visited; ri-- {
			ans = append(ans, matrix[ri][left])
			matrix[ri][left] = visited
		}

		left++
		right--
		top++
		bottom--
	}

	return ans
}

// 右下左上， 碰到尾部或者visited元素掉头，直到返回值中有m*n个元素
// @lc code=end

