/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	shortestStrLen := len(strs[0])
	for i, _ := range strs {
		if l := len(strs[i]); l < shortestStrLen {
			shortestStrLen = l
		}
	}

	sliceLen := len(strs)
	for i := 0; i < shortestStrLen; i++ {
		for j := 1; j < sliceLen; j++ {
			if strs[j][i] != strs[0][i] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0][0:shortestStrLen]
}

// @lc code=end

