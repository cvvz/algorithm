/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	subAns := []int{}

	var backtracking func(int, int, int)
	backtracking = func(i, k, n int) {
		if i > n || i > 9 {
			return
		}
		if k == 0 {
			return
		}

		backtracking(i+1, k, n)

		subAns = append(subAns, i)
		defer func() {
			subAns = subAns[:len(subAns)-1]
		}()

		if i == n && k == 1 {
			ans = append(ans, append([]int{}, subAns...))
		} else {
			backtracking(i+1, k-1, n-i)
		}
	}

	backtracking(1, k, n)

	return ans
}

// @lc code=end

