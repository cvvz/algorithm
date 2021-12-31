/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start, end := 0, len(s)-1

	for start <= end {
		if !isNumberOrChar(s[start]) {
			start++
			continue
		}
		if !isNumberOrChar(s[end]) {
			end--
			continue
		}
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func isNumberOrChar(c byte) bool{
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
// @lc code=end

