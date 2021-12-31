/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)

	for i := range ans {
		ans[i] = make([]int, i+1)
		for j:=0;j<i+1;j++ {
			if j == 0 || j==i{
				ans[i][j] = 1
			} else {
				ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
			}
		}
	}

	return ans
}

// golang 二维切片的用法，i为行，j为列
// 时间复杂度：O(n2)
// 空间复杂度：O(1)
// @lc code=end

