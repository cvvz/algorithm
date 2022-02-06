/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	maxLen := 1
	ans := s[:1]

	// 字符串长度
	for l := 2; l <= len(s); l++ {
		// 字符串begin下标
		for begin := 0; l+begin <= len(s); begin++ {
			end := l + begin - 1
			if s[begin] != s[end] {
				dp[begin][end] = false
				continue
			}
			if l == 2 || l == 3 {
				dp[begin][end] = true
			} else {
				dp[begin][end] = dp[begin+1][end-1]
			}
			if dp[begin][end] == true && l > maxLen {
				ans = s[begin : end+1]
			}
		}
	}
	return ans
}

// 🌟根据长度来进行枚举

// @lc code=end

