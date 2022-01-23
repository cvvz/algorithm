/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	ans := 0
	for i, j := 0, 0; i < len(s) && j < len(g); {
		if s[i] >= g[j] {
			ans++
			i++
			j++
		} else {
			i++
		}
	}

	return ans
}

// @lc code=end

