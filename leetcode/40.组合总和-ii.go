/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	n := len(candidates)

	// sort
	sort.Ints(candidates)

	var subResult []int

	var backtracking func(int, int)
	backtracking = func(i int, target int) {
		if i >= n {
			return
		}

		// 不带i
		// 🌟🌟🌟🌟🌟🌟去重：如果选择不带candidates[i]，那么直接跳过其他相同数字
		j := i
		for j+1 < n && candidates[j+1] == candidates[j] {
			j++
		}
		backtracking(j+1, target)

		// 带i
		subResult = append(subResult, candidates[i])
		// 🌟🌟🌟需要在回溯时把最后一个值删除
		defer func() {
			subResult = subResult[:len(subResult)-1]
		}()

		if target == candidates[i] {
			// 🌟🌟🌟🌟🌟🌟这里不能直接用subResult，否则内存会被覆盖
			result = append(result, append([]int{}, subResult...))
		} else if target < candidates[i] {
			return
		} else {
			backtracking(i+1, target-candidates[i])
		}
	}

	backtracking(0, target)

	return result
}

// @lc code=end

