/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	ans := [][]int{}
	subAns := []int{}

	var backtracking func(int, int)
	backtracking = func(start int, k int) {
		if start > n {
			return
		}
		if k > n-start+1 {
			return
		}

		backtracking(start+1, k)

		subAns = append(subAns, start)
		defer func() {
			subAns = subAns[:len(subAns)-1]
		}()

		if k == 1 {
			ans = append(ans, append([]int{}, subAns...))
		} else {
			backtracking(start+1, k-1)
		}
	}

	backtracking(1, k)

	return ans
}

// @lc code=end

