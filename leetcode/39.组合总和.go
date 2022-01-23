/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	subAns := []int{}

	n := len(candidates)

	sort.Ints(candidates)

	var backtracking func(int, int)

	backtracking = func(i, target int) {
		if i >= n {
			return
		}

		// 不选i
		backtracking(i+1, target)

		// 选i
		subAns = append(subAns, candidates[i])
		defer func() {
			subAns = subAns[:len(subAns)-1]
		}()

		if candidates[i] == target {
			// 🌟🌟🌟🌟🌟🌟这里不能直接用subResult，否则内存会被覆盖
			ans = append(ans, append([]int{}, subAns...))
		} else if target > candidates[i] {
			backtracking(i, target-candidates[i])
		}
	}

	backtracking(0, target)

	return ans
}

// @lc code=end

