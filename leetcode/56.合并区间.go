/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	ans := [][]int{intervals[0]}

	n := len(intervals)

	for i := 1; i < n; i++ {
		if intervals[i][0] > ans[len(ans)-1][1] {
			ans = append(ans, intervals[i])
		} else if intervals[i][1] > ans[len(ans)-1][1] {
			ans[len(ans)-1][1] = intervals[i][1]
		}
	}

	return ans
}

// @lc code=end

